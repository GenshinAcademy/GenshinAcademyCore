package news

import (
	"ga/internal/academy_core"
	academyModels "ga/internal/academy_core/models"
	"ga/internal/academy_core/repositories"
	"ga/internal/academy_core/repositories/find_parameters"
	"ga/internal/services/news/models"
	"ga/pkg/genshin_core/models/languages"
	gFindParameters "ga/pkg/genshin_core/repositories/find_parameters"

	"errors"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

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

// GetAllNews returns all news in specified language
func (service *Service) GetAllNews(c *gin.Context) {
	var language = languages.GetLanguage(languages.ConvertStringsToLanguages(strings.Split(c.GetHeader("Accept-Language"), ",")))

	// TODO: GetProvider should return error if provider is not found
	var newsRepo = service.core.GetProvider(language).CreateNewsRepo()

	var result = newsRepo.FindNews(find_parameters.NewsFindParameters{SliceOptions: gFindParameters.SliceParameters{Offset: uint32(c.GetUint("offset")), Limit: uint32(c.GetUint("limit"))}})

	var news []academyModels.News = result

	c.JSON(http.StatusOK,
		news)
}

func (service *Service) CreateNews(c *gin.Context) {
	// Get languages repositories
	langs := languages.ConvertStringsToLanguages(strings.Split(c.GetHeader("Accept-Language"), ","))
	newsRepos := make(map[languages.Language]repositories.INewsRepository, len(langs))
	for _, lang := range langs {
		if lang == languages.DefaultLanguage {
			continue
		}

		// TODO: GetProvider should return error if provider is not found
		repo := service.core.GetProvider(lang).CreateNewsRepo()
		newsRepos[lang] = repo
	}

	// Read request body
	var requestData models.NewsLocalized
	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if len(requestData.Title) == 0 || len(requestData.Description) == 0 || requestData.Preview == "" || requestData.Redirect == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	if requestData.Title[languages.DefaultLanguage] == "" || requestData.Description[languages.DefaultLanguage] == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Default language localization strings are required"})
		return
	}

	if requestData.Preview == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Preview provided", "message": requestData.Preview})
		return
	}

	if !requestData.Redirect.IsUrl() || requestData.Redirect == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Redirect provided", "message": requestData.Redirect})
		return
	}

	// Create general fields using default repository
	defaultRepo := service.core.GetProvider(languages.DefaultLanguage).CreateNewsRepo()
	var news academyModels.News
	news.Preview = requestData.Preview
	news.RedirectUrl = requestData.Redirect
	news.CreatedAt = time.Now()
	news.Title = requestData.Title[languages.DefaultLanguage]
	news.Description = requestData.Description[languages.DefaultLanguage]

	// Add to database
	result, err := defaultRepo.AddNews(&news)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create news"})
		return
	}

	// Update localization fields
	if len(requestData.Title) > 0 || len(requestData.Description) > 0 {
		var wg sync.WaitGroup
		errChan := make(chan error)

		for lang, repo := range newsRepos {
			wg.Add(1)
			go func(id academyModels.AcademyId, data models.NewsLocalized, repo repositories.INewsRepository, lang languages.Language) {
				defer wg.Done()
				if err := updateLocalizationFields(id, data, repo, lang); err != nil {
					errChan <- err
				}
			}(academyModels.AcademyId(result), requestData, repo, lang)
		}

		if err := <-errChan; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update localization fields"})
			return
		}
	}
	c.JSON(http.StatusOK,
		result)
}

func (service *Service) UpdateNews(c *gin.Context) {
	// Get languages repositories
	langs := languages.ConvertStringsToLanguages(strings.Split(c.GetHeader("Accept-Language"), ","))
	newsRepos := make(map[languages.Language]repositories.INewsRepository, len(langs))
	for _, lang := range langs {
		if lang == languages.DefaultLanguage {
			continue
		}

		// TODO: GetProvider should return error if provider is not found
		repo := service.core.GetProvider(lang).CreateNewsRepo()
		newsRepos[lang] = repo
	}

	// Get news ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	// Read & validate request body
	var requestData models.NewsLocalized
	err = c.BindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if len(requestData.Title) == 0 && len(requestData.Description) == 0 && requestData.Preview == "" && requestData.Redirect == "" && requestData.CreatedAt.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No update fields provided"})
		return
	}

	if requestData.Preview != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Preview provided", "message": requestData.Preview})
		return
	}

	if !requestData.Redirect.IsUrl() && requestData.Redirect != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Redirect provided", "message": requestData.Redirect})
		return
	}

	// Update general fields
	defaultRepo := service.core.GetProvider(languages.DefaultLanguage).CreateNewsRepo()
	news := defaultRepo.FindNewsById(academyModels.AcademyId(id))
	if requestData.Preview != "" {
		news.Preview = requestData.Preview
	}
	if requestData.Redirect != "" {
		news.RedirectUrl = requestData.Redirect
	}
	if !requestData.CreatedAt.IsZero() {
		news.CreatedAt = requestData.CreatedAt
	}

	// Commit to database
	if err = defaultRepo.UpdateNews(&news); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update news", "message": err.Error()})
		return
	}

	// Update localization fields
	var wg sync.WaitGroup
	errChan := make(chan error)

	for lang, repo := range newsRepos {
		wg.Add(1)
		go func(id academyModels.AcademyId, data models.NewsLocalized, repo repositories.INewsRepository, lang languages.Language) {
			defer wg.Done()
			if err := updateLocalizationFields(id, data, repo, lang); err != nil {
				errChan <- err
			}
		}(academyModels.AcademyId(id), requestData, repo, lang)
	}

	if err := <-errChan; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update news"})
		return
	}

	c.JSON(http.StatusOK, "success")
}

// TODO: Delete news

func updateLocalizationFields(id academyModels.AcademyId, requestData models.NewsLocalized, repo repositories.INewsRepository, lang languages.Language) error {
	result := repo.FindNewsById(id)
	if result == *new(academyModels.News) {
		return errors.New("news not found")
	}

	if value, ok := requestData.Title[lang]; ok {
		result.Title = value
	}

	if value, ok := requestData.Description[lang]; ok {
		result.Description = value
	}

	if err := repo.UpdateNews(&result); err != nil {
		return err
	}

	return nil
}
