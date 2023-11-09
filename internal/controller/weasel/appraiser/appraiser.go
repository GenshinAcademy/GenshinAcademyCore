package appraiser

import (
	"fmt"
	"ga/internal/controller/handlers"
	"ga/internal/models"
	"ga/internal/types"
	"ga/pkg/url"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type WeaselAppraiserCharacter struct {
	CharacterId types.CharacterId       `json:"character_id"`
	Name        string                  `json:"name"`
	Element     types.Element           `json:"element"`
	IconUrl     url.Url                 `json:"icon_url"`
	StatsProfit []models.ArtifactProfit `json:"stats_profit"`
}

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

type GetAllWeaselAppraisersResponse struct {
	Id              types.CharacterId              `json:"id"`
	Name            string                         `json:"name"`
	IconUrl         string                         `json:"icon_url"`
	ArtifactProfits types.CharacterArtifactProfits `json:"artifact_profits"`
}

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

	response := make([]GetAllWeaselAppraisersResponse, len(result))
	for i := range result {
		response[i] = GetAllWeaselAppraisersResponse{
			Id:              result[i].Id,
			Name:            result[i].Name,
			IconUrl:         result[i].IconsUrl[types.FrontFace],
			ArtifactProfits: result[i].CharacterArtifactProfits,
		}
	}

	ctx.JSON(http.StatusOK, response)
}

type UpdateCharacterStatsRequest types.CharacterArtifactProfits

func (c *Controller) UpdateStats(ctx *gin.Context) {
	id := ctx.Param("id")

	var requestData UpdateCharacterStatsRequest
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
