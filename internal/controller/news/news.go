package news

import (
	"ga/internal/controller/handlers"
	"ga/internal/models"
	"ga/internal/types"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

type GetAllNewsResponse struct {
	Id          types.NewsId `json:"id" example:"1" extensions:"x-order=1"`
	Title       string       `json:"title" example:"News title" extensions:"x-order=2"`
	Description string       `json:"description" example:"News description" extensions:"x-order=3"`
	PreviewUrl  string       `json:"preview_url" example:"https://example.com/news/preview.webp" extensions:"x-order=4"`
	RedirectUrl string       `json:"redirect_url" example:"https://example.com/news/redirect" extensions:"x-order=5"`
	CreatedAt   time.Time    `json:"created_at" example:"2020-01-01T00:00:0" extensions:"x-order=6"`
} //	@name	News

type CreateNewsRequest struct {
	Title       types.LocalizedString `json:"title,omitempty" ga:"required,localized" example:"en:News title,ru:Заголовок новости" extensions:"x-order=1"`
	Description types.LocalizedString `json:"description,omitempty" ga:"required,localized" example:"en:News description,ru:Описание новости" extensions:"x-order=2"`
	Preview     types.LocalizedString `json:"preview,omitempty" ga:"required,localized" example:"en:news-en.webp,ru:news-ru.webp" extensions:"x-order=3"`
	RedirectUrl types.LocalizedString `json:"redirect_url,omitempty" ga:"required,localized" example:"en:https://example.com/news/redirect_en,ru:https://example.com/news/redirect_ru" extensions:"x-order=4"`
} //	@name	CreateNewsRequest

type UpdateNewsRequest struct {
	Title       types.LocalizedString `json:"title,omitempty" ga:"required,localized" example:"en:News title,ru:Заголовок новости" extensions:"x-order=1"`
	Description types.LocalizedString `json:"description,omitempty" ga:"required,localized" example:"en:News description,ru:Описание новости" extensions:"x-order=2"`
	Preview     types.LocalizedString `json:"preview,omitempty" ga:"required,localized" example:"en:news-en.webp,ru:news-ru.webp" extensions:"x-order=3"`
	RedirectUrl types.LocalizedString `json:"redirect_url,omitempty" ga:"required,localized" example:"en:https://example.com/news/redirect_en,ru:https://example.com/news/redirect_ru" extensions:"x-order=4"`
} //	@name	UpdateNewsRequest

type NewsService interface {
	GetNews(language types.Language, offset int, limit int, sort string) ([]models.News, error)
	CreateNews(news *models.NewsMultilingual) error
	UpdateNews(id types.NewsId, news *models.NewsMultilingual) error
}

type Controller struct {
	newsService NewsService
}

func New(newsService NewsService) *Controller {
	return &Controller{
		newsService: newsService,
	}
}

// GetAll godoc
//
//	@Summary		Get all news
//	@Description	Get a list of news with optional offset and limit for specific language.
//	@Tags			news
//	@Produce		json
//	@Param			Accept-Languages	header	string	false	"Preferred languages for response content"	default(en,ru)
//	@Param			offset				query	int		false	"Offset for pagination"
//	@Param			limit				query	int		false	"Limit for pagination"
//	@Param			sort				query	string	false	"Sort by field"	Enums(asc,desc)	Default(desc)
//	@Success		200					{array}	GetAllNewsResponse
//	@Failure		500
//	@Router			/news [get]
func (c *Controller) GetAll(ctx *gin.Context) {
	var (
		language = handlers.GetLanguage(ctx)
		offset   = ctx.GetInt("offset")
		limit    = ctx.GetInt("limit")
		sort     = ctx.GetString("sort")
	)

	if sort == "" {
		sort = "desc"
	}

	result, err := c.newsService.GetNews(language, offset, limit, sort)
	if err != nil {
		log.Errorf("failed to get news: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get news"})
		return
	}

	response := make([]GetAllNewsResponse, len(result))
	for i := range result {
		response[i] = GetAllNewsResponse{
			Id:          result[i].Id,
			Title:       result[i].Title,
			Description: result[i].Description,
			PreviewUrl:  result[i].Preview,
			RedirectUrl: result[i].RedirectUrl,
			CreatedAt:   result[i].CreatedAt,
		}
	}

	ctx.JSON(http.StatusOK, response)
}

// Create godoc
//
//	@Summary		Create news
//	@Description	Create news with the provided data.
//	@Tags			news
//	@Accept			json
//	@Param			news	body	CreateNewsRequest	true	"News data"
//	@Security		ApiKeyAuth
//	@Success		201
//	@Failure		400,500
//	@Router			/news [post]
func (c *Controller) Create(ctx *gin.Context) {
	var requestData CreateNewsRequest
	if err := ctx.BindJSON(&requestData); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request data"})
		return
	}

	if err := handlers.HasAllFields(requestData); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	if err := handlers.HasLocalizedDefault(requestData, types.DefaultLanguage); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	newsMultilingual := &models.NewsMultilingual{
		Title:       requestData.Title,
		Description: requestData.Description,
		Preview:     requestData.Preview,
		RedirectUrl: requestData.RedirectUrl,
	}

	if err := c.newsService.CreateNews(newsMultilingual); err != nil {
		log.Errorf("failed to create news: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to create news"})
		return
	}

	ctx.Status(http.StatusCreated)
}

// Update godoc
//
//	@Summary		Update news
//	@Description	Update news with the provided id and data.
//	@Tags			news
//	@Accept			json
//	@Param			news	body	UpdateNewsRequest	true	"News data. All fields are optional but at least 1 is required."
//	@Security		ApiKeyAuth
//	@Success		202
//	@Failure		400,500
//	@Router			/news/{id} [patch]
func (c *Controller) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var requestData UpdateNewsRequest
	if err := ctx.BindJSON(&requestData); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request data"})
		return
	}

	if err := handlers.HasAnyFields(requestData); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	newsMultilingual := &models.NewsMultilingual{
		Title:       requestData.Title,
		Description: requestData.Description,
		Preview:     requestData.Preview,
		RedirectUrl: requestData.RedirectUrl,
	}

	if err := c.newsService.UpdateNews(types.NewsId(id), newsMultilingual); err != nil {
		log.Errorf("failed to update news: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to update news"})
		return
	}

	ctx.Status(http.StatusAccepted)
}
