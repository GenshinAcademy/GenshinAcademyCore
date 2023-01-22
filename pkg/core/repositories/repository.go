package repositories

import "ga/pkg/core/models"

type IRepository interface {
	GetLanguage() string
}

type ICharacterRepository interface {
	IRepository
	FindCharacterById(characterId models.ModelId) models.Character
	FindCharacters(parameters CharacterFindParameters) []models.Character
	AddCharacter(character models.Character) (models.ModelId, error)
}

type ICharacterIconRepository interface {
	IRepository
	FindIconsByCharacterId(characterId models.ModelId) []models.CharacterIcon
}

type IRepositoryProvider interface {
	GetLanguage() string
	SetLanguage(language string)
	NewCharacterRepo() ICharacterRepository
	//NewCharacterIconRepo() ICharacterIconRepository
}
