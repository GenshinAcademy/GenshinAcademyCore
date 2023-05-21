package academy

import (
	"ga/internal/academy_core/models"
	"ga/internal/academy_core/repositories"
	"ga/internal/db_postgres/cache"

	"gorm.io/gorm"
)

type PostgresAcademyProvider struct {
	gormConnection *gorm.DB
	language       *models.Language
	cache          *cache.Cache
}

func CreateAcademyProvider(connection *gorm.DB, language *models.Language, cache *cache.Cache) repositories.IRepositoryProvider {
	return PostgresAcademyProvider{
		gormConnection: connection,
		language:       language,
		cache:          cache,
	}
}

func (provider PostgresAcademyProvider) GetLanguage() *models.Language {
	return provider.language
}

func (provider PostgresAcademyProvider) NewCharacterRepo() repositories.ICharacterRepository {
	var academyRepository = CreatePostgresCharacterRepository(
		provider.gormConnection,
		provider.language,
		provider.cache)

	return academyRepository
}

func (provider PostgresAcademyProvider) CreateNewsRepo() repositories.INewsRepository {
	var academyRepository = CreatePostgresNewsRepository(
		provider.gormConnection,
		provider.language,
		provider.cache)

	return academyRepository
}

func (provider PostgresAcademyProvider) CreateTableRepo() repositories.ITableRepository {
	var academyRepository = CreatePostgresTableRepository(
		provider.gormConnection,
		provider.language,
		provider.cache)

	return academyRepository
}

func (provider PostgresAcademyProvider) CreateWeaponRepo() repositories.IWeaponRepository {
	var academyRepository = CreatePostgresWeaponRepository(
		provider.gormConnection,
		provider.language,
		provider.cache)

	return academyRepository
}
