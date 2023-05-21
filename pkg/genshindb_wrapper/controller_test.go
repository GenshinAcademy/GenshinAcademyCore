package genshindb_wrapper

import (
	"ga/pkg/genshindb_wrapper/enums"
	"testing"

	"go.uber.org/zap"
)

var (
	logger, _ = zap.NewDevelopment()
	api       = Create("https://genshin-db-api.vercel.app/api", enums.English, logger.Sugar())
)

func TestGetCharacter(t *testing.T) {
	character, err := api.GetCharacter("hutao")

	if err != nil {
		t.Errorf("GetCharacter() returned an error: %v", err)
	}

	if character.Name != "Hu Tao" {
		t.Errorf("GetCharacter() returned the wrong character: %+v", character)
	}

	_, err = api.GetCharacter("")
	if err == nil {
		t.Error("GetCharacter() did not return an error for an invalid query")
	}
}

func TestGetAllCharacters(t *testing.T) {
	characters, err := api.GetAllCharacters()

	if err != nil {
		t.Errorf("GetAllCharacters() returned an error: %v", err)
	}

	if len(characters) < 1 {
		t.Errorf("GetAllCharacters() returned the wrong data: %+v", characters)
	}
}

func TestGetAllCharactersNames(t *testing.T) {
	characters, err := api.GetAllCharactersNames()

	if err != nil {
		t.Errorf("GetAllCharactersNames() returned an error: %v", err)
	}

	if len(characters) < 1 {
		t.Errorf("GetAllCharactersNames() returned the wrong data: %+v", characters)
	}
}
