package appraiser

import (
	"fmt"
	"ga/internal/controller/handlers"
	"ga/internal/models"
	"ga/internal/types"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type WeaselAppraiserCharacter struct {
	Name            string                         `json:"name" `
	Element         types.Element                  `json:"element"`
	IconUrl         string                         `json:"icon_url"`
	ArtifactProfits types.CharacterArtifactProfits `json:"artifact_profits"`
}

type WeaselAppraiserUpdateStatsRequest types.CharacterArtifactProfits

type WeaselAppraiserCharacterDoc struct {
	Name            string                      `json:"name" example:"Hu Tao" extensions:"x-order=1"`
	Element         types.Element               `json:"element" example:"Pyro" extensions:"x-order=2"`
	IconUrl         string                      `json:"icon_url" example:"https://example.com/hu_tao.webp" extensions:"x-order=3"`
	ArtifactProfits CharacterArtifactProfitsDoc `json:"artifact_profits" extensions:"x-order=4"`
} // @name WeaselAppraiserCharacter

type CharacterArtifactProfitsDoc struct {
	SubStats models.ArtifactProfit `json:"substats" extensions:"x-order=1"`
	Flower   models.ArtifactProfit `json:"flower" extensions:"x-order=2"`
	Plume    models.ArtifactProfit `json:"plume" extensions:"x-order=3"`
	Sands    models.ArtifactProfit `json:"sands" extensions:"x-order=4"`
	Goblet   models.ArtifactProfit `json:"goblet" extensions:"x-order=5"`
	Circlet  models.ArtifactProfit `json:"circlet" extensions:"x-order=6"`
} // @name CharacterArtifactProfits

type WeaselAppraiserService interface {
	GetAll(language types.Language) ([]models.WeaselAppraiserCharacter, error)
	UpdateStats(id types.CharacterId, artifactProfits types.CharacterArtifactProfits) error
	ValidateArtifactProfitsStructure(artifactProfits types.CharacterArtifactProfits) (bool, []types.ArtifactSlot, []types.StatType)
	GetValidArtifactSlots() map[types.ArtifactSlot]bool
	GetValidArtifactTypes() map[types.StatType]bool
}

type Controller struct {
	weaselAppraiserService WeaselAppraiserService
}

func New(weaselAppraiserService WeaselAppraiserService) *Controller {
	return &Controller{
		weaselAppraiserService: weaselAppraiserService,
	}
}

// GetAll godoc
//
//	@Summary		Get all characters with stats
//	@Description	Get a list of characters with artifact profits.
//	@Tags			characters, weasel, appraiser
//	@Produce		json
//	@Param			Accept-Languages	header	string	false	"Preferred languages for response content"	default(en,ru)
//	@Success		200					{array}	WeaselAppraiserCharacterDoc
//	@Failure		500
//	@Router			/characters/stats [get]
func (c *Controller) GetAll(ctx *gin.Context) {
	var (
		language = handlers.GetLanguage(ctx)
	)

	result, err := c.weaselAppraiserService.GetAll(language)
	if err != nil {
		log.Errorf("failed to get characters: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get characters"})
		return
	}

	response := make([]WeaselAppraiserCharacter, len(result))
	for i := range result {
		response[i] = WeaselAppraiserCharacter{
			Name:            result[i].Name,
			Element:         result[i].Element,
			IconUrl:         result[i].IconsUrl[types.FrontFace],
			ArtifactProfits: result[i].CharacterArtifactProfits,
		}
	}

	ctx.JSON(http.StatusOK, response)
}

// UpdateStats godoc
//
//	@Summary		Update character stats
//	@Description	Update character stats with the provided data.
//	@Tags			characters, weasel, appraiser
//	@Accept			json
//	@Param			stats	body	CharacterArtifactProfitsDoc	true	"Character stats. All fields are optional but at least 1 is required."
//	@Security		ApiKeyAuth
//	@Success		202
//	@Failure		400,500
//	@Router			/characters/stats/{id} [patch]
func (c *Controller) UpdateStats(ctx *gin.Context) {
	id := ctx.Param("id")

	var requestData WeaselAppraiserUpdateStatsRequest
	if err := ctx.BindJSON(&requestData); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if ok, invalidSlots, invalidStatTypes := c.weaselAppraiserService.ValidateArtifactProfitsStructure(types.CharacterArtifactProfits(requestData)); !ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"invalidArtifactSlots":     fmt.Sprintf("%v", invalidSlots),
			"invalidArtifactStatTypes": fmt.Sprintf("%v", invalidStatTypes),
			"validArtifactSlots":       fmt.Sprintf("%v", c.weaselAppraiserService.GetValidArtifactSlots()),
			"validArtifactStatTypes":   fmt.Sprintf("%v", c.weaselAppraiserService.GetValidArtifactTypes()),
		})
		return
	}

	if err := c.weaselAppraiserService.UpdateStats(types.CharacterId(id), types.CharacterArtifactProfits(requestData)); err != nil {
		log.Errorf("failed to update stats: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to update stats"})
		return
	}

	ctx.Status(http.StatusAccepted)
}
