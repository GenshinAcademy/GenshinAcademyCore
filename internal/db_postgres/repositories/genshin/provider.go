package genshin

import (
    "ga/internal/db_postgres/cache"
    "ga/internal/db_postgres/repositories/academy"
    "ga/pkg/genshin_core/models/languages"
    "ga/pkg/genshin_core/repositories"
    "gorm.io/gorm"
)

type PostgresGenshinCoreProvder struct {
    gormConnection *gorm.DB
    language languages.Language
    cache *cache.Cache
}

func CreatePostgresGenshinCoreProvider(connection *gorm.DB, language languages.Language, cache *cache.Cache) PostgresGenshinCoreProvder {
    return PostgresGenshinCoreProvder{
        gormConnection: connection,
        language: language,
        cache: cache,
    }
}

func (provider PostgresGenshinCoreProvder) GetLanguage() languages.Language {
    return provider.language
}

func (provider PostgresGenshinCoreProvder) NewCharacterRepo() repositories.CharacterRepository {
    var language = academy.CreatePostresLanguageRepository(provider.gormConnection).FindLanguage(provider.GetLanguage())
    var academyRepository = academy.CreatePostgresCharacterRepository(
        provider.gormConnection,
        language,
        provider.cache)
    
    return PostgresGenshinCharacterRepository{
        academyRepo: &academyRepository,
    }
}