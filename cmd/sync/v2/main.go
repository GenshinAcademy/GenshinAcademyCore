package main

import (
	"encoding/json"
	"ga/internal/config"
	"ga/internal/db"
	"ga/internal/models"
	assets_service "ga/internal/service/assets"
	character_service "ga/internal/service/characters"
	"ga/internal/types"
	"ga/pkg/genshin_db"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const (
	dataPath = ".data"
)

var (
	cfg      *config.Config
	provider *db.Provider
)

// getCharacter retrieves the character object from a specified JSON file in a given language, by matching the character's name.
func getCharacter(name string, language genshin_db.Language) (genshin_db.Character, error) {
	var filePath = filepath.Join(".", dataPath, string(language), name)
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return genshin_db.Character{}, err
	}

	var character genshin_db.Character

	if err := json.Unmarshal(fileData, &character); err != nil {
		return genshin_db.Character{}, err
	}

	return character, nil
}

// getCharacters retrieves []genshin_db.Character slice from all JSON files in a given language.
func getCharacters(language genshin_db.Language) ([]genshin_db.Character, error) {
	var path = filepath.Join(".", dataPath, string(language))
	charactersFiles, err := os.ReadDir(path)

	if err != nil {
		return nil, err
	}

	var (
		characters   = make([]genshin_db.Character, len(charactersFiles))
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

	select {
	case err := <-characterErr:
		return nil, err
	default:
		return characters, nil
	}
}

var languages map[types.Language]genshin_db.Language

func init() {
	var err error
	cfg, err = config.New()
	if err != nil {
		panic(err)
	}

	var dbConfig = &db.PostgresConfig{
		Host:         cfg.DBHost,
		Port:         cfg.DBPort,
		Username:     cfg.DBUserName,
		Password:     cfg.DBUserPassword,
		DatabaseName: cfg.DBName,
	}

	p, err := db.NewPostgresProvider(dbConfig)
	if err != nil {
		log.Fatal(err)
	}
	provider = p

	languages = make(map[types.Language]genshin_db.Language)
	languages[types.English] = genshin_db.English
	languages[types.Russian] = genshin_db.Russian
}

func main() {
	log.WithFields(log.Fields{
		"language": languages[types.DefaultLanguage],
	}).Info("Getting default language characters from theBowja's data")

	gdbCharacters, err := getCharacters(languages[types.DefaultLanguage])
	if err != nil {
		panic(err)
	}

	characterService := character_service.New(assets_service.New(cfg.AssetsPath, cfg.AssetsHost), provider.GetCharacterRepository())

	timerStart := time.Now()

	for _, gdbCharacter := range gdbCharacters {
		log.WithFields(log.Fields{
			"character": gdbCharacter.Name,
		}).Info("Converting character from theBowja's data to Genshin Academy format")

		character := convertCharacter(gdbCharacter)

		multilingualCharacter := &models.CharacterMultilingual{
			Id:          character.Id,
			Name:        make(types.LocalizedString),
			Description: make(types.LocalizedString),
			Rarity:      character.Rarity,
			Element:     character.Element,
			WeaponType:  character.WeaponType,
			IconsUrl:    character.IconsUrl,
		}

		multilingualCharacter.Name[types.DefaultLanguage] = character.Name
		multilingualCharacter.Description[types.DefaultLanguage] = character.Description
		mlMutex := &sync.RWMutex{}

		var wg sync.WaitGroup

		for languageCode, language := range languages {
			if languageCode == types.DefaultLanguage {
				continue
			}

			wg.Add(1)
			go func(language genshin_db.Language, languageCode types.Language) {
				defer wg.Done()

				log.WithFields(log.Fields{
					"character": character.Id,
					"language":  language,
				}).Info("Adding localization info to database")

				localChar, err := getCharacter(strings.ToLower(strings.ReplaceAll(character.Name, " ", ""))+".json", language)
				if err != nil {
					panic(err)
				}
				mlMutex.Lock()
				multilingualCharacter.Name[languageCode] = localChar.Name
				multilingualCharacter.Description[languageCode] = localChar.Description
				mlMutex.Unlock()
			}(language, languageCode)
		}

		wg.Wait()
		err = characterService.CreateCharacter(multilingualCharacter)
		if err != nil {
			panic(err)
		}
	}

	timerEnd := time.Now() // Gorutine test
	log.WithFields(log.Fields{
		"time": timerEnd.Sub(timerStart).Seconds(),
	}).Info("Program has finished")
}

// convertCharacter converts character from genshin-db by theBowja to genshin-core model
func convertCharacter(input genshin_db.Character) (output models.Character) {
	output.Name = input.Name
	output.Description = input.Description

	switch input.Element {
	case genshin_db.Geo:
		output.Element = types.Geo
	case genshin_db.Dendro:
		output.Element = types.Dendro
	case genshin_db.Cryo:
		output.Element = types.Cryo
	case genshin_db.Pyro:
		output.Element = types.Pyro
	case genshin_db.Hydro:
		output.Element = types.Hydro
	case genshin_db.Electro:
		output.Element = types.Electro
	case genshin_db.Anemo:
		output.Element = types.Anemo
	default:
		output.Element = types.UndefinedElement
	}

	switch input.Rarity {
	case genshin_db.Legendary:
		output.Rarity = types.Legendary
	default:
		output.Rarity = types.Epic
	}

	switch input.WeaponType {
	case genshin_db.Sword:
		output.WeaponType = types.Sword
	case genshin_db.Bow:
		output.WeaponType = types.Bow
	case genshin_db.Claymore:
		output.WeaponType = types.Claymore
	case genshin_db.Catalyst:
		output.WeaponType = types.Catalyst
	case genshin_db.Polearm:
		output.WeaponType = types.Polearm
	default:
		output.WeaponType = types.UndefinedWeapon
	}

	return output
}
