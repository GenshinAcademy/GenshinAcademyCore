package repositories

import "ga/pkg/core/models"

type IRepository interface {
	GetLanguage() models.Language
}

type ILanguageRepository interface {
	AddLanguage(language *models.Language)
	FindLanguage(lang string) models.Language
}

type ICharacterRepository interface {
	IRepository
	GetCharacterNames(parameters CharacterFindParameters) []string
	FindCharacterById(characterId models.ModelId) models.Character
	FindCharacters(parameters CharacterFindParameters) []models.Character
	AddCharacter(character models.Character) (models.ModelId, error)
}

type ICharacterIconRepository interface {
	FindIconsByCharacterId(characterId models.ModelId) []models.CharacterIcon
}

type IRepositoryProvider interface {
	GetLanguage() string
	SetLanguage(language string)
	NewCharacterRepo() ICharacterRepository
	//NewCharacterIconRepo() ICharacterIconRepository
	NewLanguageRepo() ILanguageRepository
}
