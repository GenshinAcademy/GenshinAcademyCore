package db_repositories

import (
	"ga/pkg/core/models"
	"ga/pkg/core/repositories"

	"gorm.io/gorm"
)

// *** Repository provider ***//
type PostgresRepositoryProvider struct {
	language       string
	gormConnection *gorm.DB
}

func (provider PostgresRepositoryProvider) GetLanguage() string {
	return provider.language
}

func (provider PostgresRepositoryProvider) SetLanguage(language string) {
	provider.language = language
}

func (provider PostgresRepositoryProvider) NewCharacterRepo() repositories.ICharacterRepository {
	return PostgresCharacterRepository{
		language:       provider.language,
		gormConnection: provider.gormConnection,
	}
}

func NewRepositoryProvider(language string) repositories.IRepositoryProvider {
	return PostgresRepositoryProvider{
		language: language,
	}
}

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
