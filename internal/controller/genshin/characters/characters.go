package characters

import (
	"ga/internal/controller/handlers"
	"ga/internal/models"
	"ga/internal/types"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type GetAllCharactersResponse struct {
	Id          types.CharacterId `json:"id" example:"hu_tao" extensions:"x-order=1"`
	Name        string            `json:"name" example:"Hu Tao" extensions:"x-order=2"`
	Description string            `json:"description" example:"English description" extensions:"x-order=3"`
	Rarity      types.Rarity      `json:"rarity" example:"5" extensions:"x-order=4"`
	Element     types.Element     `json:"element" example:"Pyro" extensions:"x-order=5"`
	WeaponType  types.WeaponType  `json:"weapon_type" example:"Polearm" extensions:"x-order=6"`
	IconsUrl    types.IconsUrl    `json:"icons_url" example:"0:https://example.com/hu_tao.webp" extensions:"x-order=7"`
} //	@name	Character

type CreateCharacterRequest struct {
	Name        types.LocalizedString `json:"name" ga:"required,localized" validate:"optional" example:"en:Hu Tao,ru:Ху Тао" extensions:"x-order=1"`
	Description types.LocalizedString `json:"description" ga:"required,localized" example:"en:English description,ru:Русское описание" extensions:"x-order=2"`
	Rarity      types.Rarity          `json:"rarity" ga:"required" example:"5" extensions:"x-order=3"`
	Element     types.Element         `json:"element" ga:"required" example:"Pyro" extensions:"x-order=4"`
	WeaponType  types.WeaponType      `json:"weapon_type" ga:"required" example:"Polearm" extensions:"x-order=5"`
} //	@name	CreateCharacterRequest

type CharactersService interface {
	GetCharacters(language types.Language, offset int, limit int) ([]models.Character, error)
	CreateCharacter(character *models.CharacterMultilingual) error
	DeleteCharacter(id types.CharacterId) error
}

type Controller struct {
	charactersService CharactersService
}

func New(charactersService CharactersService) *Controller {
	return &Controller{
		charactersService: charactersService,
	}
}

// GetAll godoc
//
//	@Summary		Get all characters
//	@Description	Get a list of characters with optional offset and limit for specific language.
//	@Tags			characters
//	@Produce		json
//	@Param			offset				query	int		false	"Offset for pagination"
//	@Param			limit				query	int		false	"Limit for pagination"
//	@Param			Accept-Languages	header	string	false	"Preferred languages for response content"	default(en,ru)
//	@Success		200					{array}	GetAllCharactersResponse
//	@Failure		500
//	@Router			/characters [get]
func (c *Controller) GetAll(ctx *gin.Context) {
	var (
		offset   = ctx.GetInt("offset")
		limit    = ctx.GetInt("limit")
		language = handlers.GetLanguage(ctx)
	)

	result, err := c.charactersService.GetCharacters(language, offset, limit)
	if err != nil {
		log.Errorf("failed to get characters: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get characters"})
		return
	}

	response := make([]GetAllCharactersResponse, len(result))
	for i := range result {
		response[i] = GetAllCharactersResponse{
			Id:          result[i].Id,
			Name:        result[i].Name,
			Description: result[i].Description,
			Rarity:      result[i].Rarity,
			Element:     result[i].Element,
			WeaponType:  result[i].WeaponType,
			IconsUrl:    result[i].IconsUrl,
		}
	}

	ctx.JSON(http.StatusOK, response)
}

// Create godoc
//
//	@Summary		Create a new character
//	@Description	Create a new character with the provided data.
//	@Tags			characters
//	@Accept			json
//	@Param			character	body	CreateCharacterRequest	true	"Character data"
//	@Security		ApiKeyAuth
//	@Success		201
//	@Failure		400,500
//	@Router			/characters [post]
func (c *Controller) Create(ctx *gin.Context) {
	var requestData CreateCharacterRequest
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

	character := &models.CharacterMultilingual{
		Name:        requestData.Name,
		Description: requestData.Description,
		Rarity:      requestData.Rarity,
		Element:     requestData.Element,
		WeaponType:  requestData.WeaponType,
	}

	if err := c.charactersService.CreateCharacter(character); err != nil {
		log.Errorf("failed to create character: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to create character"})
		return
	}

	ctx.Status(http.StatusCreated)
}

// Delete godoc
//
//	@Summary		Delete character
//	@Description	Delete character by id.
//	@Tags			characters
//	@Security		ApiKeyAuth
//	@Success		200
//	@Failure		500
//	@Router			/characters/{id} [delete]
func (c *Controller) Delete(ctx *gin.Context) {
	id := types.CharacterId(ctx.Param("id"))

	if err := c.charactersService.DeleteCharacter(id); err != nil {
		log.Errorf("failed to delete character: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to delete character"})
		return
	}

	ctx.Status(http.StatusOK)
}
