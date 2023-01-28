package db_repositories

import (
	//"ga/pkg/core/models"

	"ga/pkg/core/models"
	"ga/pkg/core/repositories"

	"gorm.io/gorm"
)

// *** Repository provider ***//
type PostgresRepositoryProvider struct {
	Language       models.Language
	GormConnection *gorm.DB
}

// Gets reository language
func (provider PostgresRepositoryProvider) GetLanguage() models.Language {
	return provider.Language
}

// Creates new postgres character repository with language specified by provider
func (provider PostgresRepositoryProvider) NewCharacterRepo() repositories.ICharacterRepository {
	return PostgresCharacterRepository{
		language:       provider.Language,
		gormConnection: provider.GormConnection,
	}
}
