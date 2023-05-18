package genshin

import (
	academy_find_parameters "ga/internal/academy_core/repositories/find_parameters"
	"ga/pkg/genshin_core/errors"

	genshin_models "ga/pkg/genshin_core/models"
	"ga/pkg/genshin_core/models/languages"
	find_parameters "ga/pkg/genshin_core/repositories/find_parameters"

	"ga/internal/db_postgres/repositories/academy"
)

type PostgresGenshinCharacterRepository struct {
	academyRepo *academy.PostgresCharacterRepository
}

func CreateGenshinCharacterRepository(repo *academy.PostgresCharacterRepository) *PostgresGenshinCharacterRepository {
	var genshinRepo = new(PostgresGenshinCharacterRepository)
	genshinRepo.academyRepo = repo

	return genshinRepo
}
func (repo PostgresGenshinCharacterRepository) GetLanguage() *languages.Language {
	result := languages.Language(repo.academyRepo.GetLanguage().LanguageName)
	return &result
}

func (repo PostgresGenshinCharacterRepository) GetCharacterIds(parameters find_parameters.CharacterFindParameters) []genshin_models.ModelId {
	var academyParams = academy_find_parameters.CharacterFindParameters{
		CharacterFindParameters: parameters,
	}

	return repo.academyRepo.GetCharacterIds(academyParams)
}

func (repo PostgresGenshinCharacterRepository) FindCharacterById(characterId genshin_models.ModelId) (genshin_models.Character, error) {
	var character, found = repo.academyRepo.FindCharacterByGenshinId(characterId)
	if !found {
		return genshin_models.Character{}, errors.CharacterNotFound(characterId)
	}

	return character.Character, nil
}

func (repo PostgresGenshinCharacterRepository) FindCharacters(parameters find_parameters.CharacterFindParameters) []genshin_models.Character {
	var characters = make([]genshin_models.Character, 0)
	var academyParams = academy_find_parameters.CharacterFindParameters{
		CharacterFindParameters: parameters,
	}

	var academyCharacters = repo.academyRepo.FindCharacters(academyParams)
	for i := 0; i < len(academyCharacters); i += 1 {
		characters = append(characters, academyCharacters[i].Character)
	}

	return characters
}

func (repo PostgresGenshinCharacterRepository) AddCharacter(model genshin_models.Character) (genshin_models.Character, error) {
	if model.Id == genshin_models.DEFAULT_ID {
		return genshin_models.Character{}, errors.EmptyId()
	}

	var academyCharacter, _ = repo.academyRepo.FindCharacterByGenshinId(model.Id)
	academyCharacter.Character = model
	result, err := repo.academyRepo.AddCharacter(academyCharacter)
	if err != nil {
		return genshin_models.Character{}, err
	}

	return result.Character, nil
}

func (repo PostgresGenshinCharacterRepository) UpdateCharacter(model genshin_models.Character) (genshin_models.Character, error) {
	if model.Id == genshin_models.DEFAULT_ID {
		return genshin_models.Character{}, errors.EmptyId()
	}

	var academyCharacter, found = repo.academyRepo.FindCharacterByGenshinId(model.Id)
	if !found {
		return genshin_models.Character{}, errors.CharacterNotFound(model.Id)
	}

	academyCharacter.Character = model
	result, err := repo.academyRepo.UpdateCharacter(academyCharacter)

	if err != nil {
		return genshin_models.Character{}, err
	}

	return result.Character, nil
}
