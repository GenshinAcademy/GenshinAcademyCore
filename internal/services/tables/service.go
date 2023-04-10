package tables

import (
	"ga/internal/academy_core"
	academyModels "ga/internal/academy_core/models"
	"ga/internal/academy_core/repositories"
	"ga/internal/academy_core/repositories/find_parameters"
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

// GetAllTables returns all tables in specified language
func (service *Service) GetAllTables(c *gin.Context) {
	// TODO: GetProvider should return error if provider is not found
	var language = languages.GetLanguage(languages.ConvertStringsToLanguages(strings.Split(c.GetHeader("Accept-Language"), ",")))

	var tablesRepo = service.core.GetProvider(language).CreateTableRepo()
	var result = tablesRepo.FindTables(find_parameters.TableFindParameters{SliceOptions: gFindParameters.SliceParameters{Offset: uint32(c.GetUint("offset")), Limit: uint32(c.GetUint("limit"))}})

	var tables []academyModels.Table = result

	c.JSON(http.StatusOK,
		tables)
}

func (service *Service) CreateTable(c *gin.Context) {
	// Get languages repositories
	langs := languages.ConvertStringsToLanguages(strings.Split(c.GetHeader("Accept-Language"), ","))
	tablesRepos := make(map[languages.Language]repositories.ITableRepository, len(langs))
	for _, lang := range langs {
		if lang == languages.DefaultLanguage {
			continue
		}

		// TODO: GetProvider should return error if provider is not found
		repo := service.core.GetProvider(lang).CreateTableRepo()
		tablesRepos[lang] = repo
	}

	// Read request body
	var requestData models.TablesLocalized
	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if len(requestData.Title) == 0 || len(requestData.Description) == 0 || requestData.Icon == "" || requestData.Redirect == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	if requestData.Title[languages.DefaultLanguage] == "" || requestData.Description[languages.DefaultLanguage] == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Default language localization strings are required"})
		return
	}

	if requestData.Icon == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Icon provided", "message": requestData.Icon})
		return
	}

	if !requestData.Redirect.IsUrl() || requestData.Redirect == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Redirect provided", "message": requestData.Redirect})
		return
	}

	// Create general fields using default repository
	defaultRepo := service.core.GetProvider(languages.DefaultLanguage).CreateTableRepo()
	var table academyModels.Table
	table.Icon = requestData.Icon
	table.RedirectUrl = requestData.Redirect
	table.Title = requestData.Title[languages.DefaultLanguage]
	table.Description = requestData.Description[languages.DefaultLanguage]

	// Add to database
	result, err := defaultRepo.AddTable(&table)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create table"})
		return
	}

	// Update localization fields
	updateLocaliztionFields(c, academyModels.AcademyId(result), requestData, tablesRepos)

	c.JSON(http.StatusOK,
		result)
}

func (service *Service) UpdateTable(c *gin.Context) {
	// Get languages repositories
	langs := languages.ConvertStringsToLanguages(strings.Split(c.GetHeader("Accept-Language"), ","))
	tablesRepos := make(map[languages.Language]repositories.ITableRepository, len(langs))
	for _, lang := range langs {
		if lang == languages.DefaultLanguage {
			continue
		}

		// TODO: GetProvider should return error if provider is not found
		repo := service.core.GetProvider(lang).CreateTableRepo()
		tablesRepos[lang] = repo
	}

	// Get tables ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	// Read & validate request body
	var requestData models.TablesLocalized
	err = c.BindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if len(requestData.Title) == 0 && len(requestData.Description) == 0 && requestData.Icon == "" && requestData.Redirect == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No update fields provided"})
		return
	}

	if requestData.Icon != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid icon provided", "message": requestData.Icon})
		return
	}

	if !requestData.Redirect.IsUrl() && requestData.Redirect != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid url provided", "message": requestData.Redirect})
		return
	}

	// Update general fields
	defaultRepo := service.core.GetProvider(languages.DefaultLanguage).CreateTableRepo()
	table := defaultRepo.FindTableById(academyModels.AcademyId(id))
	if requestData.Icon != "" {
		table.Icon = requestData.Icon
	}
	if requestData.Redirect != "" {
		table.RedirectUrl = requestData.Redirect
	}

	// Commit to database
	if err = defaultRepo.UpdateTable(&table); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update table", "message": err.Error()})
		return
	}

	// Update localization fields
	if err := updateLocaliztionFields(c, academyModels.AcademyId(id), requestData, tablesRepos); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update table", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "success")
}

// TODO: Delete table

func updateLocaliztionFields(c *gin.Context, id academyModels.AcademyId, requestData models.TablesLocalized, tablesRepos map[languages.Language]repositories.ITableRepository) error {
	if len(requestData.Title) > 0 || len(requestData.Description) > 0 {
		for lang, repo := range tablesRepos {
			result := repo.FindTableById(academyModels.AcademyId(id))
			if result == *new(academyModels.Table) {
				return errors.New("table not found")
			}

			if value, ok := requestData.Title[lang]; ok {
				result.Title = value
			}

			if value, ok := requestData.Description[lang]; ok {
				result.Description = value
			}

			if err := repo.UpdateTable(&result); err != nil {
				return err
			}
		}
	}
	return nil
}
