package db_repositories

import (
	"ga/pkg/core/models"
	"ga/pkg/core/repositories"

	"gorm.io/gorm"
)

// *** Character repository ***//
type PostgresCharacterRepository struct {
	language       string
	gormConnection *gorm.DB
}

func (repo PostgresCharacterRepository) GetLanguage() string {
	return repo.language
}

func (repo PostgresCharacterRepository) FindCharacterById(characterId models.ModelId) models.Character {
	panic("TODO")
}

func (repo PostgresCharacterRepository) FindCharacters(parameters repositories.CharacterFindParameters) []models.Character {
	panic("TODO")
}

func (repo PostgresCharacterRepository) AddCharacter(character models.Character) (models.ModelId, error) {
	panic("TODO")
}

func (repo PostgresCharacterRepository) GetCharacterNames(parameters repositories.CharacterFindParameters) []string {
	panic("TODO")

}
