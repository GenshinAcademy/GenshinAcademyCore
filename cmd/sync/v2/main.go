package main

import (
	"encoding/json"
	"ga/internal/academy_core"
	academy_models "ga/internal/academy_core/models"
	"ga/internal/configuration"
	academy_postgres "ga/internal/db_postgres/implementation/academy"
	core "ga/pkg/genshin_core"
	gc_models "ga/pkg/genshin_core/models"
	gc_enums "ga/pkg/genshin_core/models/enums"
	"ga/pkg/genshin_core/models/languages"
	"ga/pkg/genshin_core/repositories"
	"ga/pkg/genshin_core/value_objects"
	gdb_enums "ga/pkg/genshindb_wrapper/enums"
	gdb_models "ga/pkg/genshindb_wrapper/models"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
)

const (
	dataPath = ".data"
)

// getCharacter retrieves the character object from a specified JSON file in a given language, by matching the character's name.
func getCharacter(name string, language string) (gdb_models.Character, error) {
	var filePath = filepath.Join(".", dataPath, language, name)
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return gdb_models.Character{}, err
	}

	var character gdb_models.Character

	if err := json.Unmarshal(fileData, &character); err != nil {
		return gdb_models.Character{}, err
	}

	return character, nil
}

// getCharacters retrieves []gdb_models.Character slice from all JSONs file in a given language.
func getCharacters(language string) ([]gdb_models.Character, error) {
	var path = filepath.Join(".", dataPath, language)
	charactersFiles, err := os.ReadDir(path)

	if err != nil {
		return nil, err
	}

	var (
		characters   = make([]gdb_models.Character, len(charactersFiles))
		characterErr = make(chan error, len(charactersFiles))
		wg           sync.WaitGroup
	)

	for i, characterFile := range charactersFiles {
		wg.Add(1)
		go func(i int, fileName string) {
			defer wg.Done()
			character, err := getCharacter(fileName, language)
			if err != nil {
				characterErr <- err
				return
			}
			characters[i] = character
		}(i, characterFile.Name())
	}

	wg.Wait()

	// Check for errors in the channel after all goroutines have completed
	select {
	case err := <-characterErr:
		return nil, err
	default:
		return characters, nil
	}
}

func main() {
	err := configuration.Init()
	if err != nil {
		panic(err)
	}

	logger := configuration.GetLogger()

	defer logger.Sync()

	var dbConfig academy_postgres.PostgresDatabaseConfiguration = academy_postgres.PostgresDatabaseConfiguration{
		Host:         configuration.ENV.DBHost,
		UserName:     configuration.ENV.DBUserName,
		UserPassword: configuration.ENV.DBUserPassword,
		DatabaseName: configuration.ENV.DBName,
		Port:         configuration.ENV.DBPort,
	}

	academy_postgres.InitializePostgresDatabase(dbConfig)

	defer academy_postgres.CleanupConnections() // Make sure to close the connection

	//Initializing gacore config and configure it for postgres db
	var config academy_core.AcademyCoreConfiguration = academy_core.AcademyCoreConfiguration{
		GenshinCoreConfiguration: core.GenshinCoreConfiguration{
			DefaultLanguage: languages.DefaultLanguage,
		},
	}

	academy_postgres.ConfigurePostgresDB(&config) // Configure postgres database

	var gacore *academy_core.AcademyCore = academy_core.CreateAcademyCore(config) //Create ga core

	var langRepo = gacore.GetLanguageRepository() // Get language repository

	// Get character repositories for all languages
	characterRepos := make(map[languages.Language]repositories.CharacterRepository, len(languages.Languages))

	for languageCode := range languages.Languages {
		var language = gacore.GetLanguageRepository().FindLanguage(languageCode)

		if language.Id == 0 {
			language = academy_models.Language{
				LanguageName: string(languageCode),
			}
			langRepo.AddLanguage(&language)
			logger.Info("Language created successfully!",
				zap.String("language", language.LanguageName))
		} else {
			logger.Info("Language found successfully!",
				zap.String("language", language.LanguageName))
		}

		characterRepos[languageCode] = gacore.AsGenshinCore().GetProvider(languageCode).NewCharacterRepo()
	}

	logger.Info("Getting default language characters from theBowja's data",
		zap.String("language", string(languages.DefaultLanguage)))

	gdbCharacters, err := getCharacters(languages.Languages[languages.DefaultLanguage])

	if err != nil {
		panic(err)
	}

	defaultCharacterRepo := characterRepos[languages.DefaultLanguage]

	timerStart := time.Now() // Gorutine test

	for _, gdbCharacter := range gdbCharacters {
		logger.Info("Coverting character from theBowja's data to Genshin Academy format",
			zap.String("character", gdbCharacter.Name))

		character := convertCharacter(gdbCharacter)

		logger.Info("Adding character to database",
			zap.String("character", character.Name))

		if err = defaultCharacterRepo.AddCharacter(&character); err != nil {
			panic(err)
		}

		var wg sync.WaitGroup // Create a WaitGroup to wait for all language updates to finish

		for languageCode := range languages.Languages {
			if languageCode == languages.DefaultLanguage {
				continue
			}

			wg.Add(1) // Increment WaitGroup counter

			go func(language languages.Language, id gc_models.ModelId) { // Launch a goroutine
				defer wg.Done() // Decrement WaitGroup counter when finished

				logger.Info("Adding localization info to database",
					zap.String("character", string(id)),
					zap.String("language", string(language)))

				// Find character in database for localization updates
				localCharFromRepo, _ := characterRepos[language].FindCharacterById(id)

				// Find character in data files by theBowja
				localChar, err := getCharacter(strings.ToLower(strings.ReplaceAll(character.Name, " ", ""))+".json", languages.Languages[language])

				if err != nil {
					panic(err)
				}

				addStrings(localChar, &localCharFromRepo)

				// Commit updates
				if err = characterRepos[language].UpdateCharacter(&localCharFromRepo); err != nil {
					panic(err)
				}
			}(languageCode, character.Id)
		}

		wg.Wait() // Wait for all language updates to finish before continuing to next character
	}

	timerEnd := time.Now() // Gorutine test
	logger.Info("Program has finished", zap.Float64("time", timerEnd.Sub(timerStart).Seconds()))
}

// convertCharacter converts character from genshin-db by theBowja to genshin-core model
func convertCharacter(input gdb_models.Character) (output gc_models.Character) {
	output.Id = gc_models.ModelId(strings.ToLower(strings.ReplaceAll(input.Name, " ", "_")))

	addStrings(input, &output)

	switch input.Element {
	case gdb_enums.Geo:
		output.Element = gc_enums.Geo
	case gdb_enums.Dendro:
		output.Element = gc_enums.Dendro
	case gdb_enums.Cryo:
		output.Element = gc_enums.Cryo
	case gdb_enums.Pyro:
		output.Element = gc_enums.Pyro
	case gdb_enums.Hydro:
		output.Element = gc_enums.Hydro
	case gdb_enums.Electro:
		output.Element = gc_enums.Electro
	case gdb_enums.Anemo:
		output.Element = gc_enums.Anemo
	default:
		output.Element = gc_enums.UndefinedElement
	}

	switch input.Rarity {
	case "5":
		output.Rarity = gc_enums.Legendary
	default:
		output.Rarity = gc_enums.Epic
	}

	switch input.Weapontype {
	case gdb_enums.Sword:
		output.Weapon = gc_enums.Sword
	case gdb_enums.Bow:
		output.Weapon = gc_enums.Bow
	case gdb_enums.Claymore:
		output.Weapon = gc_enums.Claymore
	case gdb_enums.Catalyst:
		output.Weapon = gc_enums.Catalyst
	case gdb_enums.Polearm:
		output.Weapon = gc_enums.Polearm
	default:
		output.Weapon = gc_enums.UndefinedWeapon
	}

	output.Icons = []value_objects.CharacterIcon{{Type: 0, Url: "/characters/icons/" + string(output.Id)}}
	return output
}

func addStrings(input gdb_models.Character, output *gc_models.Character) {
	output.Name = input.Name
	output.FullName = input.FullName
	output.Description = input.Description
	output.Title = input.Title
}
