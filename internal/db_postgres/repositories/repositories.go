package db_repositories

import (
	//"ga/pkg/core/models"
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
