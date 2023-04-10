package repositories

import (
	"ga/internal/academy_core/models"
	"ga/pkg/genshin_core/repositories/find_parameters"

	"gorm.io/gorm"
)

var (
	fullSaveSessionRef = &gorm.Session{
		FullSaveAssociations: true,
	}
)

type QueryBuilder struct {
	connection *gorm.DB
}

func CreateQueryBuilder(dbConnection *gorm.DB) QueryBuilder {
	return QueryBuilder{connection: dbConnection}
}

func CreateUpdateQueryBuilder(dbConnection *gorm.DB) QueryBuilder {
	return QueryBuilder{connection: dbConnection.Session(fullSaveSessionRef)}
}

func (builder QueryBuilder) GetConnection() *gorm.DB {
	return builder.connection
}

func (builder QueryBuilder) UpdateAssociations() QueryBuilder {
	return QueryBuilder{connection: builder.connection.Session(fullSaveSessionRef)}
}

func (builder QueryBuilder) PreloadAll(repo PostgresRepositoryWithStrings) QueryBuilder {
	return builder.
		Preload(repo).
		PreloadStrings(repo)
}

func (builder QueryBuilder) FilterById(repo CanSelectById, ids []models.AcademyId) QueryBuilder {
	var connect = builder.connection

	if len(ids) == 1 {
		connect = connect.Where(repo.GetIdField()+" = ?", ids[0])
	} else {
		connect = connect.Where(repo.GetIdField()+" IN ?", ids)
	}

	return QueryBuilder{connection: connect}
}

func (builder QueryBuilder) PreloadStrings(repo PostgresRepositoryWithStrings) QueryBuilder {
	var connect = builder.connection
	var strings = repo.GetStringPreloads()

	for _, preload := range strings {
		connect = connect.Preload(preload, "language_id = ?", repo.GetLanguage().Id)
	}
	return QueryBuilder{connection: connect}
}

func (builder QueryBuilder) Preload(repo PostgresRepositoryWithPreloads) QueryBuilder {
	var connect = builder.connection
	var preloads = repo.GetPreloads()

	for _, preload := range preloads {
		connect = connect.Preload(preload)
	}

	return QueryBuilder{connection: connect}
}

func (builder QueryBuilder) Slice(parameters *find_parameters.SliceParameters) QueryBuilder {
	var connect = builder.connection

	if parameters.Limit != 0 || parameters.Offset != 0 {
		connect = connect.
			Offset(int(parameters.Offset)).
			Limit(int(parameters.Limit))
	}

	return QueryBuilder{connection: connect}
}
