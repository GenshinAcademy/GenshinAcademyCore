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
		"RedirectUrl.StringValues",
	}
	tablePreloads = make([]string, 0)
)

type PostgresTableRepository struct {
	PostgresBaseRepository
}

func CreatePostgresTableRepository(connection *gorm.DB, language *academy_models.Language, cache *cache.Cache) PostgresTableRepository {
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

func (repo PostgresTableRepository) FindTableById(id academy_models.AcademyId) (academy_models.Table, error) {
	var table db_models.Table
	var ids = make([]academy_models.AcademyId, 1)
	ids[0] = id

	var connection = repositories.CreateQueryBuilder(repo.GetConnection()).
		PreloadAll(repo).
		FilterById(repo, ids).
		GetConnection()

	if err := connection.Find(&table).Error; err != nil {
		return academy_models.Table{}, err
	}

	return repo.mapper.MapTableFromDbModel(table), nil
}

func (repo PostgresTableRepository) FindTables(parameters find_parameters.TableFindParameters) ([]academy_models.Table, error) {
	var selectedTables []db_models.Table = make([]db_models.Table, 0)

	var queryBuilder = repositories.CreateQueryBuilder(repo.GetConnection()).
		PreloadAll(repo)

	if len(parameters.Ids) > 0 {
		queryBuilder = queryBuilder.FilterById(repo, parameters.Ids)
	} else {
		if parameters.SortParameters.IdSort != find_parameters.SortNone {
			queryBuilder = queryBuilder.OrderBy(genericIdField, parameters.SortParameters.IdSort)
		}

		queryBuilder = queryBuilder.Slice(&parameters.SliceOptions)
	}

	if err := queryBuilder.GetConnection().Find(&selectedTables).Error; err != nil {
		return nil, err
	}

	var resultTables = make([]academy_models.Table, len(selectedTables))
	for index, table := range selectedTables {
		resultTables[index] = repo.mapper.MapTableFromDbModel(table)
	}

	return resultTables, nil
}

func (repo PostgresTableRepository) AddTable(table academy_models.Table) (academy_models.Table, error) {
	var dbTable = repo.mapper.MapDbTableFromModel(table)

	var connection = repositories.CreateQueryBuilder(repo.GetConnection()).
		PreloadAll(repo).
		GetConnection()
	if err := connection.Create(&dbTable).Error; err != nil {
		return academy_models.Table{}, err
	}
	result := repo.mapper.MapTableFromDbModel(dbTable)

	db_postgres.GetCache().UpdateTableStrings(dbTable)

	return result, nil
}

func (repo PostgresTableRepository) UpdateTable(table academy_models.Table) (academy_models.Table, error) {
	if table.Id == academy_models.UNDEFINED_ID {
		return academy_models.Table{}, errors.New("not existing table provided")
	}

	var dbTable = repo.mapper.MapDbTableFromModel(table)

	var connection = repositories.CreateUpdateQueryBuilder(repo.GetConnection()).
		PreloadAll(repo).
		GetConnection()
	if err := connection.Save(&dbTable).Error; err != nil {
		return academy_models.Table{}, err
	}

	db_postgres.GetCache().UpdateTableStrings(dbTable)

	table.Id = academy_models.AcademyId(dbTable.Id)
	return table, nil
}
