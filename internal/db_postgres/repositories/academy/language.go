package academy

import (
	academy_models "ga/internal/academy_core/models"
	db_mappers "ga/internal/db_postgres/mappers"
	db_models "ga/internal/db_postgres/models"
	"ga/pkg/genshin_core/models/languages"

	"gorm.io/gorm"
)

// PostgresLanguageRepository Postgres language repository
type PostgresLanguageRepository struct {
	gormConnection *gorm.DB
}

func (repo PostgresLanguageRepository) GetConnection() *gorm.DB {
    return repo.gormConnection
}

// CreatePostresLanguageRepository Creates language repository with provided gorm connection
func CreatePostresLanguageRepository(connection *gorm.DB) PostgresLanguageRepository {
	return PostgresLanguageRepository{
		gormConnection: connection,
	}
}

// AddLanguage Adds language
func (repo PostgresLanguageRepository) AddLanguage(language *academy_models.Language) {
	var langModel = repo.FindLanguage(languages.Language(language.LanguageName))
	if langModel.Id != academy_models.UNDEFINED_ID {
		panic("Language with this name already exists")
	}

	var langDbModel = db_models.Language{
		Name: language.LanguageName,
	}

	repo.gormConnection.Create(&langDbModel)
	//TODO: Error
	language.Id = academy_models.AcademyId(langDbModel.Id)
}

// FindLanguage Finds language by name
func (repo PostgresLanguageRepository) FindLanguage(lang languages.Language) academy_models.Language {
	var langDbModel = db_models.Language{}

	repo.gormConnection.Where("name = ?", &lang).First(&langDbModel)
	//TODO: Error
	return db_mappers.Mapper{}.MapLanguageFromDbModel(&langDbModel)
}
