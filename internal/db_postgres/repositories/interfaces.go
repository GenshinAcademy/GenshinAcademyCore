package repositories

import (
    "ga/internal/academy_core/models"
    "ga/internal/db_postgres/cache"
    "gorm.io/gorm"
)

//TODO: Add docs

type PostgresRepository interface {
    GetConnection() *gorm.DB
}

type PostgresRepositoryWithLanguage interface {
    PostgresRepository
    GetLanguage() models.Language
}

type PostgresRepositoryWithCache interface {
    PostgresRepository
    GetCache() *cache.Cache
}

type PostgresRepositoryWithPreloads interface {
    PostgresRepository
    GetPreloads() []string
}

type PostgresRepositoryWithStrings interface {
    PostgresRepositoryWithPreloads
    PostgresRepositoryWithLanguage
    GetStringPreloads() []string
}

type CanSelectById interface {
    PostgresRepository
    GetIdField() string
}
