package academy

import (
    "ga/internal/academy_core/models"
    db_mappers "ga/internal/db_postgres/mappers"
    "ga/pkg/genshin_core/repositories/find_parameters"
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

func PreloadAll(repo PostgresRepositoryWithStrings) *gorm.DB {
    var resultConnection = PreloadStrings(repo)
    resultConnection = Preload(repo)

    return resultConnection
}

func FilterById(repo CanSelectById, ids []models.AcademyId) *gorm.DB {
    var connection = repo.GetConnection()

    if len(ids) == 1 {
        connection = connection.Where(repo.GetIdField() + " = ?", ids[0])
        return connection
    }

    connection = connection.Where(repo.GetIdField() + " IN ?", ids)

    return connection
}

func PreloadStrings(repo PostgresRepositoryWithStrings) *gorm.DB {
    var resultConnection = repo.GetConnection()
    var strings = repo.GetStringPreloads()

    for _, preload := range strings {
        resultConnection = resultConnection.Preload(preload, "language_id = ?", repo.GetLanguage().Id)
    }
    return resultConnection
}

func Preload(repo PostgresRepositoryWithPreloads) *gorm.DB {
    var resultConnection = repo.GetConnection()
    var preloads = repo.GetPreloads()

    for _, preload := range preloads {
        resultConnection = resultConnection.Preload(preload)
    }

    return resultConnection
}

func Slice(connection *gorm.DB, parameters *find_parameters.SliceParameters) *gorm.DB {
    if parameters.Amount != 0 || parameters.Offset != 0 {
        connection = connection.
            Offset(int(parameters.Offset)).
            Limit(int(parameters.Amount))
    }

    return connection
}
