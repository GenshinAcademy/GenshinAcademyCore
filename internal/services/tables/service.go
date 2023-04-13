package tables

import (
	"ga/internal/academy_core"
	academyModels "ga/internal/academy_core/models"
	"ga/internal/academy_core/repositories"
	"ga/internal/academy_core/repositories/find_parameters"
	"ga/internal/services/handlers"
	"ga/internal/services/tables/models"
	"ga/pkg/genshin_core/models/languages"
	gFindParameters "ga/pkg/genshin_core/repositories/find_parameters"

	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Service struct {
	core *academy_core.AcademyCore
}

func CreateService(core *academy_core.AcademyCore) *Service {
	var result *Service = new(Service)
	result.core = core
	return result
}

// GetAll returns all tables in specified language
func (service *Service) GetAll(c *gin.Context) {
	// TODO: GetProvider should return error if provider is not found
	var language = languages.GetLanguage(languages.ConvertStringsToLanguages(strings.Split(c.GetHeader("Accept-Language"), ",")))

	var tablesRepo = service.core.GetProvider(language).CreateTableRepo()
	var result = tablesRepo.FindTables(
		find_parameters.TableFindParameters{
			SliceOptions: gFindParameters.SliceParameters{
				Offset: uint32(c.GetUint("offset")),
				Limit:  uint32(c.GetUint("limit"))}})

	var tables []academyModels.Table = result

	c.JSON(http.StatusOK,
		tables)
}

func (service *Service) Create(c *gin.Context) {
	// Get languages repositories
	langs := languages.ConvertStringsToLanguages(strings.Split(c.GetHeader("Accept-Language"), ","))
	tablesRepos := make(map[languages.Language]repositories.ITableRepository, len(langs))
	for _, lang := range langs {
		if lang == languages.DefaultLanguage {
			continue
		}

		// TODO: GetProvider should return error if provider is not found
		repo := service.core.GetProvider(&lang).CreateTableRepo()
		tablesRepos[lang] = repo
	}

	// Read request body
	var requestData models.TablesLocalized
	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request data"})
		return
	}

	if err := handlers.HasAllFields(requestData); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := handlers.HasLocalizedDefault(requestData, languages.DefaultLanguage); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// Create general fields using default repository
	defaultRepo := service.core.GetProvider(&languages.DefaultLanguage).CreateTableRepo()
	var table academyModels.Table
	table.Icon = requestData.Icon
	table.RedirectUrl = requestData.Redirect
	table.Title = requestData.Title[languages.DefaultLanguage]
	table.Description = requestData.Description[languages.DefaultLanguage]

	// Add to database
	var results []academyModels.Table
	result, err := defaultRepo.AddTable(&table)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create table", "message": err.Error()})
		return
	}

	results = append(results, *result)

	// Update localization fields
	if len(requestData.Title) > 0 || len(requestData.Description) > 0 {
		errChan := make(chan error)

		for lang, repo := range tablesRepos {
			go func(id academyModels.AcademyId, data models.TablesLocalized, repo repositories.ITableRepository, lang languages.Language) {
				result, err := updateLocalizationFields(id, data, repo, lang)
				if err != nil {
					errChan <- err
				}
				results = append(results, *result)
				errChan <- nil
			}(result.Id, requestData, repo, lang)
		}
		for range tablesRepos {
			if err := <-errChan; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update localization fields", "message": err.Error()})
				return
			}
		}
	}

	c.JSON(http.StatusOK, results)
}

func (service *Service) Update(c *gin.Context) {
	// Get languages repositories
	langs := languages.ConvertStringsToLanguages(strings.Split(c.GetHeader("Accept-Language"), ","))
	tablesRepos := make(map[languages.Language]repositories.ITableRepository, len(langs))
	for _, lang := range langs {
		if lang == languages.DefaultLanguage {
			continue
		}

		// TODO: GetProvider should return error if provider is not found
		repo := service.core.GetProvider(&lang).CreateTableRepo()
		tablesRepos[lang] = repo
	}

	// Get tables ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	// Read & validate request body
	var requestData models.TablesLocalized
	err = c.BindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request data"})
		return
	}

	if err := handlers.HasAnyFields(requestData); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// Update general fields
	defaultRepo := service.core.GetProvider(&languages.DefaultLanguage).CreateTableRepo()
	table := defaultRepo.FindTableById(academyModels.AcademyId(id))
	if requestData.Icon != "" {
		table.Icon = requestData.Icon
	}
	if requestData.Redirect != "" {
		table.RedirectUrl = requestData.Redirect
	}

	if value, ok := requestData.Title[languages.DefaultLanguage]; ok {
		table.Title = value
	}

	if value, ok := requestData.Description[languages.DefaultLanguage]; ok {
		table.Description = value
	}

	// Commit to database
	var results []academyModels.Table
	result, err := defaultRepo.UpdateTable(table)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update table", "message": err.Error()})
		return
	}
	results = append(results, *result)

	// Update localization fields
	if len(requestData.Title) > 0 || len(requestData.Description) > 0 {
		errChan := make(chan error, len(tablesRepos))

		for lang, repo := range tablesRepos {
			go func(id academyModels.AcademyId, data models.TablesLocalized, repo repositories.ITableRepository, lang languages.Language) {
				result, err := updateLocalizationFields(id, data, repo, lang)
				if err != nil {
					errChan <- err
				}
				results = append(results, *result)
				errChan <- nil
			}(academyModels.AcademyId(id), requestData, repo, lang)
		}
		for range tablesRepos {
			if err := <-errChan; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update localization fields", "message": err.Error()})
				return
			}
		}
	}

	c.JSON(http.StatusOK, results)
}

// TODO: Delete table

func updateLocalizationFields(id academyModels.AcademyId, requestData models.TablesLocalized, repo repositories.ITableRepository, lang languages.Language) (*academyModels.Table, error) {
	result := repo.FindTableById(id)
	if result == nil {
		return nil, errors.New("table not found")
	}

	if value, ok := requestData.Title[lang]; ok {
		result.Title = value
	}

	if value, ok := requestData.Description[lang]; ok {
		result.Description = value
	}

	newResult, err := repo.UpdateTable(result)
	if err != nil {
		return nil, err
	}

	return newResult, nil
}
