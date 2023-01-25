package db_repositories

import (
	//"ga/pkg/core/models"
	db_mappers "ga/internal/db_postgres/mappers"
	db_models "ga/internal/db_postgres/models"
	"ga/pkg/core/models"
	"ga/pkg/core/repositories"

	"gorm.io/gorm"
)

// *** Repository provider ***//
type PostgresRepositoryProvider struct {
	language       string
	GormConnection *gorm.DB
}

func (provider PostgresRepositoryProvider) GetLanguage() string {
	return provider.language
}

func (provider PostgresRepositoryProvider) SetLanguage(language string) {
	provider.language = language
}

func (provider PostgresRepositoryProvider) NewCharacterRepo() repositories.ICharacterRepository {

	var lang db_models.Db_Language
	provider.GormConnection.Where("name = ?", provider.GetLanguage()).First(lang)

	var langModel models.Language = db_mappers.LanguageFromDbModel(&lang)
	return PostgresCharacterRepository{
		language:       langModel,
		gormConnection: provider.GormConnection,
	}
}

func (provider PostgresRepositoryProvider) NewLanguageRepo() repositories.ILanguageRepository {
	return PostgresLanguageRepository{
		gormConnection: provider.GormConnection,
	}
}

func NewRepositoryProvider(language string) repositories.IRepositoryProvider {
	return PostgresRepositoryProvider{
		language: language,
	}
}
