package news

import (
	"errors"
	"ga/internal/academy_core"
	"ga/internal/academy_core/models"
	"ga/internal/academy_core/repositories"
	"ga/internal/academy_core/repositories/find_parameters"
	url "ga/internal/academy_core/value_objects/url"
	"ga/pkg/genshin_core/models/languages"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type NewsService struct {
	core *academy_core.AcademyCore
}

func CreateService(core *academy_core.AcademyCore) *NewsService {
	var result *NewsService = new(NewsService)
	result.core = core
	return result
}

// GetAllNews returns all news in specified language
func (service *NewsService) GetAllNews(c *gin.Context) {
	// TODO: GetProvider should return error if provider is not found
	var language = languages.GetLanguage(languages.ConvertStringsToLanguages(strings.Split(c.GetHeader("Accept-Language"), ",")))

	var newsRepo = service.core.GetProvider(language).CreateNewsRepo()
	var result = newsRepo.FindNews(find_parameters.NewsFindParameters{})

	var news []models.News = result

	c.JSON(http.StatusOK,
		news)
}

func (service *NewsService) CreateNews(c *gin.Context) {
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
	var requestData newsJson
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

	if !requestData.Preview.IsUrl() || requestData.Preview == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid url provided", "message": requestData.Preview})
		return
	}

	if !requestData.Redirect.IsUrl() || requestData.Redirect == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid url provided", "message": requestData.Redirect})
		return
	}

	// Create general fields using default repository
	defaultRepo := service.core.GetProvider(languages.DefaultLanguage).CreateNewsRepo()
	var news models.News
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
	updateLocaliztionFields(c, models.AcademyId(result), requestData, newsRepos)

	c.JSON(http.StatusOK,
		result)
}

func (service *NewsService) UpdateNews(c *gin.Context) {
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
	var requestData newsJson
	err = c.BindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if len(requestData.Title) == 0 && len(requestData.Description) == 0 && requestData.Preview == "" && requestData.Redirect == "" && requestData.CreatedAt.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No update fields provided"})
		return
	}

	if !requestData.Preview.IsUrl() && requestData.Preview != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid url provided", "message": requestData.Preview})
		return
	}

	if !requestData.Redirect.IsUrl() && requestData.Redirect != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid url provided", "message": requestData.Redirect})
		return
	}

	// Update general fields
	defaultRepo := service.core.GetProvider(languages.DefaultLanguage).CreateNewsRepo()
	news := defaultRepo.FindNewsById(models.AcademyId(id))
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
	if err := updateLocaliztionFields(c, models.AcademyId(id), requestData, newsRepos); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update news", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "success")
}

// TODO: Delete news

func updateLocaliztionFields(c *gin.Context, id models.AcademyId, requestData newsJson, newsRepos map[languages.Language]repositories.INewsRepository) error {
	if len(requestData.Title) > 0 || len(requestData.Description) > 0 {
		for lang, repo := range newsRepos {
			result := repo.FindNewsById(models.AcademyId(id))
			if result == *new(models.News) {
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
		}
	}
	return nil
}

type newsJson struct {
	Id          models.AcademyId              `json:"id,omitempty"`
	Title       map[languages.Language]string `json:"title,omitempty"`
	Description map[languages.Language]string `json:"description,omitempty"`
	Preview     url.Url                       `json:"preview,omitempty"`
	Redirect    url.Url                       `json:"redirect,omitempty"`
	CreatedAt   time.Time                     `json:"created_at,omitempty"`
}
