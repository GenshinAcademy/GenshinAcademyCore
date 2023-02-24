package academy

import (
    "ga/internal/academy_core/models"
    "ga/internal/academy_core/repositories"
    "ga/internal/db_postgres/cache"
    "gorm.io/gorm"
)

type PostgresAcademyProvider struct {
    gormConnection *gorm.DB
    language models.Language
    cache *cache.Cache
}

func CreateAcademyProvider(connection *gorm.DB, language models.Language, cache *cache.Cache) PostgresAcademyProvider{
    return PostgresAcademyProvider{
        gormConnection: connection,
        language: language,
        cache: cache,
    }
}

func (provider PostgresAcademyProvider) GetLanguage() models.Language {
    return provider.language
}

func (provider PostgresAcademyProvider) NewCharacterRepo() repositories.ICharacterRepository {
    var academyRepository = CreatePostgresCharacterRepository(
        provider.gormConnection,
        provider.language,
        provider.cache)

    return academyRepository
}