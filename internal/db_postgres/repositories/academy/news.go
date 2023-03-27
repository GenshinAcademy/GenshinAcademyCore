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
	newsStringPreloads = []string{
		"Title.StringValues",
		"Description.StringValues",
	}

	newsPreloads = make([]string, 0)
)

type PostgresNewsRepository struct {
	PostgresBaseRepository
}

func CreatePostgresNewsRepository(connection *gorm.DB, language academy_models.Language, cache *cache.Cache) PostgresNewsRepository {
	return PostgresNewsRepository{
		PostgresBaseRepository: PostgresBaseRepository{
			language:       language,
			gormConnection: connection,
			mapper:         db_mappers.CreateMapper(language.LanguageName, language, cache),
		},
	}
}

func (repo PostgresNewsRepository) GetIdField() string {
	return genericIdField
}

func (repo PostgresNewsRepository) GetStringPreloads() []string {
	return newsStringPreloads
}

func (repo PostgresNewsRepository) GetPreloads() []string {
	return newsPreloads
}

func (repo PostgresNewsRepository) FindNewsById(id academy_models.AcademyId) academy_models.News {
	var selectedNews db_models.News
	var ids = make([]academy_models.AcademyId, 1)
	ids[0] = id

    var connection = repositories.CreateQueryBuilder(repo.GetConnection()).
        PreloadAll(repo).
        FilterById(repo, ids).
        GetConnection()

	connection.Find(&selectedNews)

	return repo.mapper.MapNewsFromDbModel(&selectedNews)
}

func (repo PostgresNewsRepository) FindNews(parameters find_parameters.NewsFindParameters) []academy_models.News {
	var selectedNews []db_models.News = make([]db_models.News, 0)

    var queryBuilder = repositories.CreateQueryBuilder(repo.GetConnection()).PreloadAll(repo)

	if len(parameters.Ids) > 0 {
		queryBuilder = queryBuilder.FilterById(repo, parameters.Ids)
	} else {
		queryBuilder = ApplyFindParameters(queryBuilder, &parameters)
	}

	queryBuilder.GetConnection().Find(&selectedNews)

	var resultNews = make([]academy_models.News, len(selectedNews))
	for index, news := range selectedNews {
		resultNews[index] = repo.mapper.MapNewsFromDbModel(&news)
	}

	return resultNews
}

func (repo PostgresNewsRepository) AddNews(news *academy_models.News) (academy_models.AcademyId, error) {
	if news == nil {
		return academy_models.UNDEFINED_ID, errors.New("null value provided")
	}

	var dbNews = repo.mapper.MapDbNewsFromModel(news)

    var connection = repositories.CreateQueryBuilder(repo.GetConnection()).
        PreloadAll(repo).
        GetConnection()
	connection.Create(&dbNews)

	db_postgres.GetCache().UpdateNewsStrings(&dbNews)

	return academy_models.AcademyId(dbNews.Id), nil
}

func (repo PostgresNewsRepository) UpdateNews(news *academy_models.News) error {
	if news == nil {
		return errors.New("null value provided")
	}
	if news.Id == academy_models.UNDEFINED_ID {
		return errors.New("not existing news provided")
	}

	var dbNews = repo.mapper.MapDbNewsFromModel(news)

    var connection = repositories.CreateQueryBuilder(repo.GetConnection()).
        PreloadAll(repo).
        GetConnection()
	connection.Save(&dbNews)

	db_postgres.GetCache().UpdateNewsStrings(&dbNews)

	return nil
}

func ApplyFindParameters(builder repositories.QueryBuilder, parameters *find_parameters.NewsFindParameters) repositories.QueryBuilder {
	if parameters.PublishTimeFrom != nil {
		//TODO
	}

	if parameters.PublishTimeTo != nil {
		//TODO
	}

	if parameters.SortByDescendingTime {
		//TODO
	}

	return builder.Slice(&parameters.SliceOptions)
}