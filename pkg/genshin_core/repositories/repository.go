package repositories

import (
	"ga/pkg/genshin_core/models"
	"ga/pkg/genshin_core/models/languages"
	"ga/pkg/genshin_core/repositories/find_parameters"
	"ga/pkg/genshin_core/value_objects"
)

type Repository interface {
	GetLanguage() *languages.Language
}

type CharacterRepository interface {
	Repository
	GetCharacterIds(parameters find_parameters.CharacterFindParameters) []models.ModelId
	FindCharacterById(characterId models.ModelId) (models.Character, error)
	FindCharacters(parameters find_parameters.CharacterFindParameters) []models.Character
	AddCharacter(character *models.Character) error
	UpdateCharacter(character *models.Character) error
}

type CharacterIconRepository interface {
	FindIconsByCharacterId(characterId models.ModelId) []value_objects.CharacterIcon
}

type RepositoryProvider interface {
	GetLanguage() *languages.Language
	NewCharacterRepo() CharacterRepository
	//NewCharacterIconRepo() ICharacterIconRepository
}
