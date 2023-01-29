package repositories

import (
	"ga/pkg/genshin_core/models"
	"ga/pkg/genshin_core/repositories/find_parameters"
)

type IRepository interface {
	GetLanguage() models.Language
}

type ILanguageRepository interface {
	AddLanguage(language *models.Language)
	FindLanguage(lang string) models.Language
}

type ICharacterRepository interface {
	IRepository
	GetCharacterIds(parameters find_parameters.CharacterFindParameters) []string
	FindCharacterById(characterId models.ModelId) models.Character
	FindCharacters(parameters find_parameters.CharacterFindParameters) []models.Character
	AddCharacter(character *models.Character) (models.ModelId, error)
	UpdateCharacter(character *models.Character)
}

type ICharacterIconRepository interface {
	FindIconsByCharacterId(characterId models.ModelId) []models.CharacterIcon
}

type IRepositoryProvider interface {
	GetLanguage() models.Language
	NewCharacterRepo() ICharacterRepository
	//NewCharacterIconRepo() ICharacterIconRepository
}
