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
	gdb_models "ga/pkg/genshindb_wrapper/models"
	"os"
	"path/filepath"
	"strings"
)

const (
	dataPath = ".data"
)

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

func getCharacters(language string) ([]gdb_models.Character, error) {
	var path = filepath.Join(".", dataPath, language)
	charactersFiles, err := os.ReadDir(path)

	if err != nil {
		return nil, err
	}

	var characters = make([]gdb_models.Character, 0, len(charactersFiles))

	for _, characterFile := range charactersFiles {

		character, err := getCharacter(characterFile.Name(), language)

		if err != nil {
			return nil, err
		}

		characters = append(characters, character)
	}

	return characters, nil
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
	defer academy_postgres.CleanupConnections()

	//Initializing gacore config and configure it for postgres db
	var config academy_core.AcademyCoreConfiguration = academy_core.AcademyCoreConfiguration{
		GenshinCoreConfiguration: core.GenshinCoreConfiguration{
			DefaultLanguage: languages.DefaultLanguage,
		},
	}
	academy_postgres.ConfigurePostgresDB(&config)

	//Create ga core
	var gacore *academy_core.AcademyCore = academy_core.CreateAcademyCore(config)

	var langRepo = gacore.GetLanguageRepository()

	//Create enLanguage if it does not exist

	// Get character repositories for all languages
	characterRepos := make(map[languages.Language]repositories.CharacterRepository, len(languages.Languages))

	for languageCode := range languages.Languages {
		var language = gacore.GetLanguageRepository().FindLanguage(languageCode)

		if language.Id == 0 {
			language = academy_models.Language{
				LanguageName: string(languageCode),
			}
			langRepo.AddLanguage(&language)
			logger.Sugar().Infow("Language created successfully!",
				"language", language)
		} else {
			logger.Sugar().Infow("Language found successfully!",
				"language", language)
		}

		characterRepos[languageCode] = gacore.AsGenshinCore().GetProvider(languageCode).NewCharacterRepo()
	}

	logger.Sugar().Infow("Getting default language characters from theBowja's data",
		"language", languages.DefaultLanguage)

	gdbCharacters, err := getCharacters(languages.Languages[languages.DefaultLanguage])

	if err != nil {
		panic(err)
	}

	defaultCharacterRepo := characterRepos[languages.DefaultLanguage]

	for _, gdbCharacter := range gdbCharacters {
		logger.Sugar().Infow("Coverting character from theBowja's data to Genshin Academy format",
			"character", gdbCharacter.Name)

		character := convertCharacter(gdbCharacter)

		logger.Sugar().Infow("Adding character to database",
			"character", character.Id)

		if err = defaultCharacterRepo.AddCharacter(&character); err != nil {
			panic(err)
		}

		for languageCode := range languages.Languages {
			if languageCode == languages.DefaultLanguage {
				continue
			}

			logger.Sugar().Infow("Adding localization info to database",
				"character", character.Id,
				"language", languageCode)

			localCharFromRepo, _ := characterRepos[languageCode].FindCharacterById(character.Id)

			localChar, err := getCharacter(strings.ToLower(strings.ReplaceAll(character.Name, " ", ""))+".json", languages.Languages[languageCode])

			if err != nil {
				panic(err)
			}

			addStrings(localChar, &localCharFromRepo)

			if err = characterRepos[languageCode].UpdateCharacter(&localCharFromRepo); err != nil {
				panic(err)
			}
		}
	}
}

// convertCharacter converts character from genshin-db by theBowja to genshin-core model
func convertCharacter(input gdb_models.Character) (output gc_models.Character) {
	output.Id = gc_models.ModelId(strings.ToLower(strings.ReplaceAll(input.Name, " ", "_")))

	addStrings(input, &output)

	switch input.Element {
	case "Geo":
		output.Element = gc_enums.Geo
	case "Dendro":
		output.Element = gc_enums.Dendro
	case "Cryo":
		output.Element = gc_enums.Cryo
	case "Pyro":
		output.Element = gc_enums.Pyro
	case "Hydro":
		output.Element = gc_enums.Hydro
	case "Electro":
		output.Element = gc_enums.Electro
	case "Anemo":
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
	case "Sword":
		output.Weapon = gc_enums.Sword
	case "Bow":
		output.Weapon = gc_enums.Bow
	case "Claymore":
		output.Weapon = gc_enums.Claymore
	case "Catalyst":
		output.Weapon = gc_enums.Catalyst
	case "Polearm":
		output.Weapon = gc_enums.Polearm
	default:
		output.Weapon = gc_enums.UndefinedWeapon
	}

	output.Icons = []value_objects.CharacterIcon{{Type: 0, Url: input.Images.Icon}}
	return output
}

func addStrings(input gdb_models.Character, output *gc_models.Character) {
	output.Name = input.Name
	output.FullName = input.FullName
	output.Description = input.Description
	output.Title = input.Title
}
