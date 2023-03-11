package genshin

import (
	"ga/internal/academy_core"
	"ga/pkg/genshin_core/models/languages"
	"ga/pkg/genshin_core/repositories/find_parameters"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GenshinService struct {
	core academy_core.AcademyCore
}

func CreateService(core academy_core.AcademyCore) *GenshinService {
	var result *GenshinService = new(GenshinService)
	return result
}

// GetAllCharacters returns all characters raw
// Requires Accept-Language header in request
func (service *GenshinService) GetAllCharacters(c *gin.Context) {
	// TODO: GetProvider should return error if provider is not found
	var characterRepo = service.core.AsGenshinCore().GetProvider(languages.Language(c.GetHeader("Accept-Language"))).NewCharacterRepo()
	var result = characterRepo.FindCharacters(find_parameters.CharacterFindParameters{})

	c.JSON(http.StatusOK,
		result)
}
