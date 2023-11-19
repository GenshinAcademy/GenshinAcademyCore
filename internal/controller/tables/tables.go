package tables

import (
	"ga/internal/controller/handlers"
	"ga/internal/models"
	"ga/internal/types"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type GetAllTablesResponse struct {
	Id          types.TableId `json:"id" example:"1" extensions:"x-order=1"`
	Title       string        `json:"title" example:"en:Table title,ru:Заголовок таблицы" extensions:"x-order=2"`
	Description string        `json:"description" example:"en:Table description,ru:Описание таблицы" extensions:"x-order=3"`
	IconUrl     string        `json:"icon_url" example:"https://example.com/tables/icon.png" extensions:"x-order=4"`
	RedirectUrl string        `json:"redirect_url" example:"https://example.com/tables/redirect" extensions:"x-order=5"`
} //	@name	Table

type CreateTableRequest struct {
	Title       types.LocalizedString `json:"title,omitempty" ga:"required,localized" example:"en:Monsters' resistances and shields,ru:Сопротивления и щиты монстров" extensions:"x-order=1"`
	Description types.LocalizedString `json:"description,omitempty" ga:"required,localized" example:"en:Elemental resistances and gauges of their elemental shields or structures along with notes on their specific mechanics that change those values.,ru:Лиза" extensions:"x-order=2"`
	Icon        string                `json:"icon,omitempty" ga:"required" example:"https://example.com/tables/shield.webp" extensions:"x-order=3"`
	RedirectUrl types.LocalizedString `json:"redirect_url,omitempty" ga:"required,localized" example:"en:https://example.com/tables/redirect_en,ru:https://example.com/tables/redirect_ru" extensions:"x-order=4"`
} //	@name	CreateTableRequest

type UpdateTableRequest struct {
	Title       types.LocalizedString `json:"title,omitempty" ga:"required,localized" example:"en:Monsters' resistances and shields,ru:Сопротивления и щиты монстров" extensions:"x-order=1"`
	Description types.LocalizedString `json:"description,omitempty" ga:"required,localized" example:"en:Elemental resistances and gauges of their elemental shields or structures along with notes on their specific mechanics that change those values.,ru:Лиза" extensions:"x-order=2"`
	Icon        string                `json:"icon,omitempty" ga:"required" example:"https://example.com/tables/shield.webp" extensions:"x-order=3"`
	RedirectUrl types.LocalizedString `json:"redirect_url,omitempty" ga:"required,localized" example:"en:https://example.com/tables/redirect_en,ru:https://example.com/tables/redirect_ru" extensions:"x-order=4"`
} //	@name	UpdateTableRequest

type TablesService interface {
	GetTables(language types.Language, offset int, limit int, sort string) ([]models.Table, error)
	CreateTable(table *models.TableMultilingual) error
	UpdateTable(id types.TableId, table *models.TableMultilingual) error
}

type Controller struct {
	tablesService TablesService
}

func New(tablesService TablesService) *Controller {
	return &Controller{
		tablesService: tablesService,
	}
}

// GetAll godoc
//
//	@Summary		Get all tables
//	@Description	Get all tables.
//	@Tags			tables
//	@Produces		json
//	@Param			Accept-Languages	header	string	false	"Preferred languages for response content"	default(en,ru)
//	@Param			offset				query	int		false	"Offset for pagination"
//	@Param			limit				query	int		false	"Limit for pagination"
//	@Param			sort				query	string	false	"Sort by field"	Enums(asc,desc)	Default(desc)
//	@Success		200					{array}	GetAllTablesResponse
//	@Failure		500
//	@Router			/tables [get]
func (c *Controller) GetAll(ctx *gin.Context) {
	var (
		language = handlers.GetLanguage(ctx)
		offset   = ctx.GetInt("offset")
		limit    = ctx.GetInt("limit")
		sort     = ctx.GetString("sort")
	)

	result, err := c.tablesService.GetTables(language, offset, limit, sort)
	if err != nil {
		log.Errorf("failed to get tables: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get tables"})
		return
	}

	response := make([]GetAllTablesResponse, len(result))
	for i := range result {
		response[i] = GetAllTablesResponse{
			Id:          result[i].Id,
			Title:       result[i].Title,
			Description: result[i].Description,
			IconUrl:     result[i].IconUrl,
			RedirectUrl: result[i].RedirectUrl,
		}
	}

	ctx.JSON(http.StatusOK, response)
}

// Create godoc
//
//	@Summary		Create table
//	@Description	Create table.
//	@Tags			tables
//	@Accept			json
//	@Param			table	body	CreateTableRequest	true	"Table data."
//	@Security		ApiKeyAuth
//	@Success		201
//	@Failure		400,500
//	@Router			/tables [post]
func (c *Controller) Create(ctx *gin.Context) {
	var requestData CreateTableRequest
	if err := ctx.BindJSON(&requestData); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "invalid request data",
			"message": err.Error(),
		})
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

	multilingualTable := &models.TableMultilingual{
		Title:       requestData.Title,
		Description: requestData.Description,
		IconUrl:     requestData.Icon,
		RedirectUrl: requestData.RedirectUrl,
	}

	if err := c.tablesService.CreateTable(multilingualTable); err != nil {
		log.Errorf("failed to create table: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to create table"})
		return
	}

	ctx.Status(http.StatusCreated)
}

// Update godoc
//
//	@Summary		Update table
//	@Description	Update table.
//	@Tags			tables
//	@Accept			json
//	@Param			table	body	UpdateTableRequest	true	"Table data. All fields are optional but at least 1 is required."
//	@Security		ApiKeyAuth
//	@Failure		400,500
//	@Router			/tables/{id} [put]
func (c *Controller) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "invalid id",
			"message": err.Error(),
		})
		return
	}

	var requestData UpdateTableRequest
	err = ctx.BindJSON(&requestData)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "invalid request data",
			"message": err.Error(),
		})
		return
	}

	if err := handlers.HasAnyFields(requestData); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	multilingualTable := &models.TableMultilingual{
		Title:       requestData.Title,
		Description: requestData.Description,
		IconUrl:     requestData.Icon,
		RedirectUrl: requestData.RedirectUrl,
	}

	if err := c.tablesService.UpdateTable(types.TableId(id), multilingualTable); err != nil {
		log.Errorf("failed to update table: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to update table"})
		return
	}

	ctx.Status(http.StatusAccepted)
}
