package academy

import (
    "ga/internal/academy_core/models"
    db_mappers "ga/internal/db_postgres/mappers"
    "gorm.io/gorm"
)

//TODO: Add docs

const(
    genericIdField string = "id"
)

type PostgresBaseRepository struct {
    gormConnection *gorm.DB
    language models.Language
    mapper db_mappers.Mapper
}

func(repo PostgresBaseRepository)  GetConnection() *gorm.DB {
    return repo.gormConnection
}

func (repo PostgresBaseRepository) GetLanguage() models.Language {
    return repo.language
}