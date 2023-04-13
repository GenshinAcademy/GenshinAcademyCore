package news

import (
	"ga/internal/academy_core"
	academyModels "ga/internal/academy_core/models"
	"ga/internal/academy_core/repositories"
	"ga/internal/academy_core/repositories/find_parameters"
	"ga/internal/services/handlers"
	"ga/internal/services/news/models"
	"ga/pkg/genshin_core/models/languages"
	gFindParameters "ga/pkg/genshin_core/repositories/find_parameters"

	"errors"
	"net/http"
	"strconv"
	"strings"
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

// GetAll returns all news in specified language
func (service *Service) GetAll(c *gin.Context) {
	var language = languages.GetLanguage(languages.ConvertStringsToLanguages(strings.Split(c.GetHeader("Accept-Language"), ",")))

	// TODO: GetProvider should return error if provider is not found
	var newsRepo = service.core.GetProvider(language).CreateNewsRepo()

	var result = newsRepo.FindNews(
		find_parameters.NewsFindParameters{
			SliceOptions: gFindParameters.SliceParameters{
				Offset: uint32(c.GetUint("offset")),
				Limit:  uint32(c.GetUint("limit"))}})

	var news []academyModels.News = result

	c.JSON(http.StatusOK,
		news)
}

func (service *Service) Create(c *gin.Context) {
	// Get languages repositories
	langs := languages.ConvertStringsToLanguages(strings.Split(c.GetHeader("Accept-Language"), ","))
	newsRepos := make(map[languages.Language]repositories.INewsRepository, len(langs))
	for _, lang := range langs {
		if lang == languages.DefaultLanguage {
			continue
		}

		// TODO: GetProvider should return error if provider is not found
		repo := service.core.GetProvider(&lang).CreateNewsRepo()
		newsRepos[lang] = repo
	}

	// Read request body
	var requestData models.NewsLocalized
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
	defaultRepo := service.core.GetProvider(&languages.DefaultLanguage).CreateNewsRepo()
	var news academyModels.News
	news.Preview = requestData.Preview
	news.RedirectUrl = requestData.Redirect
	news.CreatedAt = time.Now()
	news.Title = requestData.Title[languages.DefaultLanguage]
	news.Description = requestData.Description[languages.DefaultLanguage]

	// Add to database
	var results []academyModels.News
	result, err := defaultRepo.AddNews(&news)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create news", "message": err.Error()})
		return
	}

	results = append(results, *result)

	// Update localization fields
	if len(requestData.Title) > 0 || len(requestData.Description) > 0 {
		errChan := make(chan error)
		for lang, repo := range newsRepos {
			go func(id academyModels.AcademyId, data models.NewsLocalized, repo repositories.INewsRepository, lang languages.Language) {

				res, err := updateLocalizationFields(id, data, repo, lang)
				if err != nil {
					errChan <- err
				}
				results = append(results, *res)
				errChan <- nil
			}(result.Id, requestData, repo, lang)
		}
		for range newsRepos {
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
	newsRepos := make(map[languages.Language]repositories.INewsRepository, len(langs))
	for _, lang := range langs {
		if lang == languages.DefaultLanguage {
			continue
		}

		// TODO: GetProvider should return error if provider is not found
		repo := service.core.GetProvider(&lang).CreateNewsRepo()
		newsRepos[lang] = repo
	}

	// Get news ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	// Read & validate request body
	var requestData models.NewsLocalized
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
	defaultRepo := service.core.GetProvider(&languages.DefaultLanguage).CreateNewsRepo()
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
	if value, ok := requestData.Title[languages.DefaultLanguage]; ok {
		news.Title = value
	}
	if value, ok := requestData.Description[languages.DefaultLanguage]; ok {
		news.Description = value
	}

	// Commit to database
	var results []academyModels.News
	result, err := defaultRepo.UpdateNews(news)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update news", "message": err.Error()})
		return
	}

	results = append(results, *result)

	// Update localization fields
	if len(requestData.Title) > 0 || len(requestData.Description) > 0 {
		errChan := make(chan error, len(newsRepos))

		for lang, repo := range newsRepos {
			go func(id academyModels.AcademyId, data models.NewsLocalized, repo repositories.INewsRepository, lang languages.Language) {
				res, err := updateLocalizationFields(id, data, repo, lang)
				if err != nil {
					errChan <- err
				}
				results = append(results, *res)
				errChan <- nil
			}(academyModels.AcademyId(id), requestData, repo, lang)
		}

		for range newsRepos {
			if err := <-errChan; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update news", "message": err.Error()})
				return
			}
		}
	}

	c.JSON(http.StatusOK, results)
}

// TODO: Delete news

func updateLocalizationFields(id academyModels.AcademyId, requestData models.NewsLocalized, repo repositories.INewsRepository, lang languages.Language) (*academyModels.News, error) {
	// TODO: Error handling
	// BUG: panic: runtime error: invalid memory address or nil pointer dereference
	// [signal 0xc0000005 code=0x0 addr=0x18 pc=0x30d5bf]
	result := repo.FindNewsById(id)
	if result == nil {
		return nil, errors.New("news not found")
	}

	if value, ok := requestData.Title[lang]; ok {
		result.Title = value
	}

	if value, ok := requestData.Description[lang]; ok {
		result.Description = value
	}

	newResult, err := repo.UpdateNews(result)
	if err != nil {
		return nil, err
	}

	return newResult, nil
}
