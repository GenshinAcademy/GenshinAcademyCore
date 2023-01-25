package db_repositories

import (
	db_mappers "ga/internal/db_postgres/mappers"
	db_models "ga/internal/db_postgres/models"
	"ga/pkg/core/models"

	"gorm.io/gorm"
)

type PostgresLanguageRepository struct {
	gormConnection *gorm.DB
}

func (repo PostgresLanguageRepository) AddLanguage(language *models.Language) {
	var langModel = db_models.Db_Language{
		Name: language.LanguageName,
	}

	repo.gormConnection.Create(&langModel)
	//TODO: Error
	language.Id = models.ModelId(langModel.Id)
}

func (repo PostgresLanguageRepository) FindLanguage(lang string) models.Language {
	var langModel = db_models.Db_Language{}

	repo.gormConnection.Where("name = ?", &lang).First(&langModel)
	//TODO: Error
	return db_mappers.LanguageFromDbModel(&langModel)
}
