package academy

import (
	models "ga/internal/academy_core/models"
	db_mappers "ga/internal/db_postgres/mappers"
	db_models "ga/internal/db_postgres/models"

	"gorm.io/gorm"
)

// Postgres language repository
type PostgresLanguageRepository struct {
	gormConnection *gorm.DB
}

// Creates language repository with provided gorm connection
func CreatePostresLanguageRepository(connection *gorm.DB) PostgresLanguageRepository {
	return PostgresLanguageRepository{
		gormConnection: connection,
	}
}

// Adds language
func (repo PostgresLanguageRepository) AddLanguage(language *models.Language) {
	var langModel = repo.FindLanguage(language.LanguageName)
	if langModel.Id != 0 {
		panic("Language with this name already exists")
	}

	var langDbModel = db_models.DbLanguage{
		Name: language.LanguageName,
	}

	repo.gormConnection.Create(&langDbModel)
	//TODO: Error
	language.Id = models.AcademyId(langDbModel.Id)
}

// Finds language by name
func (repo PostgresLanguageRepository) FindLanguage(lang string) models.Language {
	var langDbModel = db_models.DbLanguage{}

	repo.gormConnection.Where("name = ?", &lang).First(&langDbModel)
	//TODO: Error
    return db_mappers.Mapper{}.LanguageFromDbModel(&langDbModel)
}