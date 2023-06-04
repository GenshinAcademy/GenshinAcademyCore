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

	"net/http"
	"net/url"
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

// GetAllTables godoc
// @Summary Get all tables from database
// @Tags tables
// @Description Retrieves all tables.
// @Produce json
// @Param Accept-Languages header string true "Result language" default(en)
// @Param offset query int false "Offset for pagination"
// @Param limit query int false "Limit for pagination"
// @Success 200 {array} academyModels.Table
// @Failure 404 {error} error "error"
// @Router /tables [get]
func (service *Service) GetAll(c *gin.Context) {
	// TODO: GetProvider should return error if provider is not found
	var language = languages.GetLanguage(languages.ConvertStringsToLanguages(strings.Split(c.GetHeader("Accept-Languages"), ",")))

	var tablesRepo = service.core.GetProvider(language).CreateTableRepo()
	var result, err = tablesRepo.FindTables(
		find_parameters.TableFindParameters{
			SliceOptions: gFindParameters.SliceParameters{
				Offset: uint32(c.GetUint("offset")),
				Limit:  uint32(c.GetUint("limit"))}})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "tables not found",
			"message": err.Error(),
		})
	}

	// Add assets path to non URL values
	const tablesPath = "tables/"

	for i := range result {
		table := &result[i]
		if !isURL(table.Icon) && table.Icon != "" {
			iconPath := tablesPath + table.Icon
			iconURL, err := service.core.GetAssetPath(iconPath)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   "failed to get asset path",
					"message": err.Error(),
				})
				return
			}
			table.Icon = string(iconURL)
		}
	}

	c.JSON(http.StatusOK,
		result)
}

func isURL(input string) bool {
	u, err := url.Parse(input)
	return err == nil && u.Scheme != ""
}

// CreateTable godoc
// @Summary Create table
// @Tags tables
// @Description Creates a new table in database.
// @Accept json
// @Produce json
// @Param Accept-Languages header string true "Languages splitted by comma. Specify each language you are adding in json body" default(en,ru)
// @Param table body models.TablesLocalized true "Table data"
// @Security ApiKeyAuth
// @Router /tables [post]
// @Success 200 {array} academyModels.Table
// @Failure 400 {string} string "error"
// @Failure 500 {object} string "error"
func (service *Service) Create(c *gin.Context) {
	// Get languages repositories
	langs := languages.ConvertStringsToLanguages(strings.Split(c.GetHeader("Accept-Languages"), ","))
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
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid request data",
			"message": err.Error(),
		})
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
	table.RedirectUrl = requestData.Redirect[languages.DefaultLanguage]
	table.Title = requestData.Title[languages.DefaultLanguage]
	table.Description = requestData.Description[languages.DefaultLanguage]

	// Add to database
	var results []academyModels.Table
	result, err := defaultRepo.AddTable(table)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "failed to create table",
			"message": err.Error(),
		})
		return
	}

	results = append(results, result)

	// Update localization fields
	if len(requestData.Title) > 0 || len(requestData.Description) > 0 {
		errChan := make(chan error)

		for lang, repo := range tablesRepos {
			go func(id academyModels.AcademyId, data models.TablesLocalized, repo repositories.ITableRepository, lang languages.Language) {
				result, err := updateLocalizationFields(id, data, repo, lang)
				if err != nil {
					errChan <- err
				}
				results = append(results, result)
				errChan <- nil
			}(result.Id, requestData, repo, lang)
		}
		for range tablesRepos {
			if err := <-errChan; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   "failed to update localization",
					"message": err.Error(),
				})
				return
			}
		}
	}

	c.JSON(http.StatusOK, results)
}

// UpdateTable godoc
// @Summary Update table
// @Tags tables
// @Description Updates selected table in database.
// @Accept json
// @Produce json
// @Param Accept-Languages header string true "Languages splitted by comma. Specify each language you are adding in json body" default(en,ru)
// @Param id path int true "Table ID"
// @Param table body models.TablesLocalized true "Table data"
// @Security ApiKeyAuth
// @Router /tables/{id} [patch]
// @Success 200 {array} academyModels.Table
// @Failure 400 {string} string "error"
// @Failure 404 {object} string "error"
// @Failure 500 {object} string "error"
func (service *Service) Update(c *gin.Context) {
	// Get languages repositories
	langs := languages.ConvertStringsToLanguages(strings.Split(c.GetHeader("Accept-Languages"), ","))
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
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid id",
			"message": err.Error(),
		})
		return
	}

	// Read & validate request body
	var requestData models.TablesLocalized
	err = c.BindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid request data",
			"message": err.Error(),
		})
		return
	}

	if err := handlers.HasAnyFields(requestData); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// Update general fields
	defaultRepo := service.core.GetProvider(&languages.DefaultLanguage).CreateTableRepo()
	table, err := defaultRepo.FindTableById(academyModels.AcademyId(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "tables not found",
			"message": err.Error(),
		})
	}

	if value, ok := requestData.Title[languages.DefaultLanguage]; ok {
		table.Title = value
	}

	if value, ok := requestData.Description[languages.DefaultLanguage]; ok {
		table.Description = value
	}
	if requestData.Icon != "" {
		table.Icon = requestData.Icon
	}
	if value, ok := requestData.Redirect[languages.DefaultLanguage]; ok {
		table.RedirectUrl = value
	}

	// Commit to database
	var results []academyModels.Table
	result, err := defaultRepo.UpdateTable(table)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "failed to update table",
			"message": err.Error(),
		})
		return
	}
	results = append(results, result)

	// Update localization fields
	if len(requestData.Title) > 0 || len(requestData.Description) > 0 {
		errChan := make(chan error, len(tablesRepos))

		for lang, repo := range tablesRepos {
			go func(id academyModels.AcademyId, data models.TablesLocalized, repo repositories.ITableRepository, lang languages.Language) {
				result, err := updateLocalizationFields(id, data, repo, lang)
				if err != nil {
					errChan <- err
				}
				results = append(results, result)
				errChan <- nil
			}(academyModels.AcademyId(id), requestData, repo, lang)
		}
		for range tablesRepos {
			if err := <-errChan; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   "failed to update localization fields",
					"message": err.Error(),
				})
				return
			}
		}
	}

	c.JSON(http.StatusOK, results)
}

// TODO: Delete table

func updateLocalizationFields(id academyModels.AcademyId, requestData models.TablesLocalized, repo repositories.ITableRepository, lang languages.Language) (academyModels.Table, error) {
	result, err := repo.FindTableById(id)
	if err != nil {
		return academyModels.Table{}, err
	}

	if value, ok := requestData.Title[lang]; ok {
		result.Title = value
	}

	if value, ok := requestData.Description[lang]; ok {
		result.Description = value
	}

	if value, ok := requestData.Redirect[lang]; ok {
		result.RedirectUrl = value
	}

	newResult, err := repo.UpdateTable(result)
	if err != nil {
		return academyModels.Table{}, err
	}

	return newResult, nil
}
