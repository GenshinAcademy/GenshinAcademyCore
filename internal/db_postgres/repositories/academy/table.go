package academy

import (
	"errors"
	academy_models "ga/internal/academy_core/models"
	"ga/internal/academy_core/repositories/find_parameters"
	"ga/internal/db_postgres"
	"ga/internal/db_postgres/cache"
	db_mappers "ga/internal/db_postgres/mappers"
	db_models "ga/internal/db_postgres/models"
	"ga/internal/db_postgres/repositories"
	"gorm.io/gorm"
)

var (
	tableStringPreloads = []string{
		"Title.StringValues",
		"Description.StringValues",
	}
	tablePreloads = make([]string, 0)
)

type PostgresTableRepository struct {
	PostgresBaseRepository
}

func CreatePostgresTableRepository(connection *gorm.DB, language academy_models.Language, cache *cache.Cache) PostgresTableRepository {
	return PostgresTableRepository{
		PostgresBaseRepository: PostgresBaseRepository{
			language:       language,
			gormConnection: connection,
			mapper:         db_mappers.CreateMapper(language.LanguageName, language, cache),
		},
	}
}

func (repo PostgresTableRepository) GetIdField() string {
	return genericIdField
}

func (repo PostgresTableRepository) GetStringPreloads() []string {
	return tableStringPreloads
}

func (repo PostgresTableRepository) GetPreloads() []string {
	return tablePreloads
}

func (repo PostgresTableRepository) FindTableById(id academy_models.AcademyId) academy_models.Table {
	var table db_models.Table
	var ids = make([]academy_models.AcademyId, 1)
	ids[0] = id

	var connection = repositories.CreateQueryBuilder(repo.GetConnection()).
		PreloadAll(repo).
		FilterById(repo, ids).
		GetConnection()

	connection.Find(&table)

	return repo.mapper.MapTableFromDbModel(&table)
}

func (repo PostgresTableRepository) FindTables(parameters find_parameters.TableFindParameters) []academy_models.Table {
	var selectedTables []db_models.Table = make([]db_models.Table, 0)

	var queryBuilder = repositories.CreateQueryBuilder(repo.GetConnection()).
		PreloadAll(repo)

	if len(parameters.Ids) > 0 {
		queryBuilder = queryBuilder.FilterById(repo, parameters.Ids)
	} else {
		queryBuilder = queryBuilder.Slice(&parameters.SliceOptions)
	}

	queryBuilder.GetConnection().Find(&selectedTables)

	var resultTables = make([]academy_models.Table, len(selectedTables))
	for index, table := range selectedTables {
		resultTables[index] = repo.mapper.MapTableFromDbModel(&table)
	}

	return resultTables
}

func (repo PostgresTableRepository) AddTable(table *academy_models.Table) (academy_models.AcademyId, error) {
	if table == nil {
		return academy_models.UNDEFINED_ID, errors.New("null value provided")
	}

	var dbTable = repo.mapper.MapDbTableFromModel(table)

	var connection = repositories.CreateQueryBuilder(repo.GetConnection()).
		PreloadAll(repo).
		GetConnection()
	connection.Create(&dbTable)

	db_postgres.GetCache().UpdateTableStrings(&dbTable)

	return academy_models.AcademyId(dbTable.Id), nil
}

func (repo PostgresTableRepository) UpdateTable(table *academy_models.Table) error {
	if table == nil {
		return errors.New("null value provided")
	}
	if table.Id == academy_models.UNDEFINED_ID {
		return errors.New("not existing table provided")
	}

	var dbTable = repo.mapper.MapDbTableFromModel(table)

	var connection = repositories.CreateUpdateQueryBuilder(repo.GetConnection()).
		PreloadAll(repo).
		GetConnection()
	connection.Save(&dbTable)

	db_postgres.GetCache().UpdateTableStrings(&dbTable)

	return nil
}
