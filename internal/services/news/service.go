package news

import (
	"fmt"
	"ga/internal/academy_core"
	academyModels "ga/internal/academy_core/models"
	"ga/internal/academy_core/repositories"
	"ga/internal/academy_core/repositories/find_parameters"
	"ga/internal/services/handlers"
	"ga/internal/services/news/models"
	"ga/pkg/genshin_core/models/languages"
	gFindParameters "ga/pkg/genshin_core/repositories/find_parameters"

	"net/http"
	"net/url"
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

// GetAllNews godoc
// @Summary Get all news
// @Tags news
// @Description Retrieves all news from database sorted by date.
// @Produce json
// @Param Accept-Languages header string true "Result language" default(en)
// @Param offset query int false "Offset for pagination"
// @Param limit query int false "Limit for pagination"
// @Success 200 {array} academyModels.News
// @Failure 404 {error} error "error"
// @Router /news [get]
func (service *Service) GetAll(c *gin.Context) {
	var language = languages.GetLanguage(languages.ConvertStringsToLanguages(strings.Split(c.GetHeader("Accept-Languages"), ",")))

	// TODO: GetProvider should return error if provider is not found
	var newsRepo = service.core.GetProvider(language).CreateNewsRepo()

	var result, err = newsRepo.FindNews(
		find_parameters.NewsFindParameters{
			SortOptions: find_parameters.NewsSortParameters{CreatedTimeSort: find_parameters.SortByDescending},
			SliceOptions: gFindParameters.SliceParameters{
				Offset: uint32(c.GetUint("offset")),
				Limit:  uint32(c.GetUint("limit"))}})

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
	}

	// Add assets path to non URL values
	const newsPath = "news/"

	for i := range result {
		news := &result[i]
		if !isURL(news.Preview) && news.Preview != "" {
			iconPath := newsPath + news.Preview
			iconURL, err := service.core.GetAssetPath(iconPath)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   "failed to get asset path",
					"message": err.Error(),
				})
				return
			}
			news.Preview = string(iconURL)
		}
	}

	for _, news := range result {
		_, err := url.Parse(news.Preview)
		if err != nil {
			newsPath := "news"

			iconUrl, err := service.core.GetAssetPath(fmt.Sprintf("%s/%s", newsPath, news.Preview))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   "failed to get assets path",
					"message": err.Error(),
				})
			}

			news.Preview = string(iconUrl)
		}
	}

	c.JSON(http.StatusOK,
		result)
}

// CreateNews godoc
// @Summary Create news
// @Tags news
// @Description Creates news in database.
// @Accept json
// @Produce json
// @Param Accept-Languages header string true "Languages splitted by comma. Specify each language you are adding in json body" default(en,ru)
// @Param news body models.NewsLocalized true "News data"
// @Security ApiKeyAuth
// @Router /news [post]
// @Success 200 {array} academyModels.News
// @Failure 400 {string} string "error"
// @Failure 500 {object} string "error"
func (service *Service) Create(c *gin.Context) {
	// Get languages repositories
	langs := languages.ConvertStringsToLanguages(strings.Split(c.GetHeader("Accept-Languages"), ","))
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
	news.Title = requestData.Title[languages.DefaultLanguage]
	news.Description = requestData.Description[languages.DefaultLanguage]
	news.Preview = requestData.Preview[languages.DefaultLanguage]
	news.RedirectUrl = requestData.Redirect[languages.DefaultLanguage]
	news.CreatedAt = time.Now()

	// Add to database
	var results []academyModels.News
	result, err := defaultRepo.AddNews(news)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create news", "message": err.Error()})
		return
	}

	results = append(results, result)

	// Update localization fields
	if len(requestData.Title) > 0 || len(requestData.Description) > 0 {
		errChan := make(chan error)
		for lang, repo := range newsRepos {
			go func(id academyModels.AcademyId, data models.NewsLocalized, repo repositories.INewsRepository, lang languages.Language) {

				res, err := updateLocalizationFields(id, data, repo, lang)
				if err != nil {
					errChan <- err
				}
				results = append(results, res)
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

// UpdateNews godoc
// @Summary Update news
// @Tags news
// @Description Updates selected news in database.
// @Accept json
// @Produce json
// @Param Accept-Languages header string true "Languages splitted by comma. Specify each language you are adding in json body" default(en,ru)
// @Param id path int true "News ID"
// @Param news body models.NewsLocalized true "News data"
// @Security ApiKeyAuth
// @Router /news/{id} [patch]
// @Success 200 {array} academyModels.News
// @Failure 400 {string} string "error"
// @Failure 404 {string} string "error"
// @Failure 500 {object} string "error"
func (service *Service) Update(c *gin.Context) {
	// Get languages repositories
	langs := languages.ConvertStringsToLanguages(strings.Split(c.GetHeader("Accept-Languages"), ","))
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
	news, err := defaultRepo.FindNewsById(academyModels.AcademyId(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
	}

	// if !requestData.CreatedAt.IsZero() {
	// 	news.CreatedAt = requestData.CreatedAt
	// }
	if value, ok := requestData.Title[languages.DefaultLanguage]; ok {
		news.Title = value
	}
	if value, ok := requestData.Description[languages.DefaultLanguage]; ok {
		news.Description = value
	}
	if value, ok := requestData.Preview[languages.DefaultLanguage]; ok {
		news.Preview = value
	}
	if value, ok := requestData.Redirect[languages.DefaultLanguage]; ok {
		news.RedirectUrl = value
	}

	// Commit to database
	var results []academyModels.News
	result, err := defaultRepo.UpdateNews(news)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update news", "message": err.Error()})
		return
	}

	results = append(results, result)

	// Update localization fields
	if len(requestData.Title) > 0 || len(requestData.Description) > 0 {
		errChan := make(chan error, len(newsRepos))

		for lang, repo := range newsRepos {
			go func(id academyModels.AcademyId, data models.NewsLocalized, repo repositories.INewsRepository, lang languages.Language) {
				res, err := updateLocalizationFields(id, data, repo, lang)
				if err != nil {
					errChan <- err
				}
				results = append(results, res)
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

func updateLocalizationFields(id academyModels.AcademyId, requestData models.NewsLocalized, repo repositories.INewsRepository, lang languages.Language) (academyModels.News, error) {
	// TODO: Error handling
	// BUG: panic: runtime error: invalid memory address or nil pointer dereference
	// [signal 0xc0000005 code=0x0 addr=0x18 pc=0x30d5bf]
	result, err := repo.FindNewsById(id)
	if err != nil {
		return academyModels.News{}, err
	}

	if value, ok := requestData.Title[lang]; ok {
		result.Title = value
	}

	if value, ok := requestData.Description[lang]; ok {
		result.Description = value
	}

	if value, ok := requestData.Preview[lang]; ok {
		result.Preview = value
	}

	if value, ok := requestData.Redirect[lang]; ok {
		result.RedirectUrl = value
	}

	newResult, err := repo.UpdateNews(result)
	if err != nil {
		return academyModels.News{}, err
	}

	return newResult, nil
}

func isURL(input string) bool {
	u, err := url.Parse(input)
	return err == nil && u.Scheme != ""
}
