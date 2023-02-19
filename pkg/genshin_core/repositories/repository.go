package repositories

import (
	"ga/pkg/genshin_core/languages"
	"ga/pkg/genshin_core/models"
	"ga/pkg/genshin_core/repositories/find_parameters"
)

type Repository interface {
	GetLanguage() languages.Language
}

type CharacterRepository interface {
	Repository
	GetCharacterIds(parameters find_parameters.CharacterFindParameters) []models.ModelId
	FindCharacterById(characterId models.ModelId) models.Character
	FindCharacters(parameters find_parameters.CharacterFindParameters) []models.Character
	AddCharacter(character *models.Character) (models.ModelId, error)
	UpdateCharacter(character *models.Character)
}

type CharacterIconRepository interface {
	FindIconsByCharacterId(characterId models.ModelId) []models.CharacterIcon
}

type RepositoryProvider interface {
	GetLanguage() languages.Language
	NewCharacterRepo() CharacterRepository
	//NewCharacterIconRepo() ICharacterIconRepository
}
