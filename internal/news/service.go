package news

import (
	"ga/internal/academy_core"
	"ga/internal/academy_core/models"
	"ga/internal/academy_core/repositories/find_parameters"
	"ga/pkg/genshin_core/models/languages"
	"net/http"
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
// Requires Accept-Language header in request
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
	// TODO: GetProvider should return error if provider is not found
	var language = languages.GetLanguage(languages.ConvertStringsToLanguages(strings.Split(c.GetHeader("Accept-Language"), ",")))

	var newsRepo = service.core.GetProvider(language).CreateNewsRepo()
	var news = &models.News{
		AcademyModel: models.AcademyModel{Id: 1},
		Title:        "Test",
		Description:  "Test Description",
		Preview:      "https://cdn.com/testPreview.webp",
		RedirectUrl:  "testRedirect.com",
		CreatedAt:    time.Now(),
	}

	var result, err = newsRepo.AddNews(news)

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			err)
	}

	c.JSON(http.StatusOK,
		result)
}
