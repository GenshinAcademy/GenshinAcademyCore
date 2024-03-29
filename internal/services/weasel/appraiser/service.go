package appraiser

import (
	"ga/internal/academy_core"
	"ga/internal/academy_core/repositories/find_parameters"
	"ga/internal/academy_core/value_objects/artifact_profit"
	"ga/internal/services/handlers"
	"ga/internal/services/weasel/appraiser/models"
	gc_models "ga/pkg/genshin_core/models"
	"ga/pkg/genshin_core/models/languages"
	"net/http"
	"strings"

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

// GetAll returns all characters profit information in specified language
// Requires Accept-Languages header in request
func (service *Service) GetAll(c *gin.Context) {
	var language = languages.GetLanguage(languages.ConvertStringsToLanguages(strings.Split(c.GetHeader("Accept-Languages"), ",")))

	// TODO: GetProvider should return error if provider is not found
	var characterRepo = service.core.GetProvider(language).NewCharacterRepo()
	var result, err = characterRepo.FindCharacters(find_parameters.CharacterFindParameters{})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	var characters []models.WeaselAppraiserCharacter
	for _, char := range result {
		if len(char.Profits) == 0 {
			continue
		}

		character, err := service.mapCharacter(char)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err})
			return
		}

		characters = append(characters, character)
	}

	c.JSON(http.StatusOK,
		characters)
}

func (service *Service) UpdateStats(c *gin.Context) {
	id := c.Param("id")
	var characterRepo = service.core.GetProvider(&languages.DefaultLanguage).NewCharacterRepo()

	var requestData []artifact_profit.ArtifactProfit
	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	for _, data := range requestData {
		if err := handlers.HasAllFields(data); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
	}

	char, err := characterRepo.FindCharacterByGenshinId(gc_models.ModelId(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
	}

	char.Profits = requestData

	result, err := characterRepo.UpdateCharacter(char)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to update character", "message": err.Error()})
	}

	c.JSON(http.StatusOK,
		result)
}
