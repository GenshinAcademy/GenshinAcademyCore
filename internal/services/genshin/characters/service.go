package characters

import (
	"ga/internal/academy_core"
	"ga/pkg/genshin_core/models/languages"
	"ga/pkg/genshin_core/repositories/find_parameters"
	gFindParameters "ga/pkg/genshin_core/repositories/find_parameters"

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

// GetAllCharacters returns all characters raw
// Requires Accept-Language header in request
func (service *Service) GetAllCharacters(c *gin.Context) {
	// TODO: GetProvider should return error if provider is not found
	var language = languages.GetLanguage(languages.ConvertStringsToLanguages(strings.Split(c.GetHeader("Accept-Language"), ",")))

	var characterRepo = service.core.AsGenshinCore().GetProvider(language).NewCharacterRepo()
	var result = characterRepo.FindCharacters(
		find_parameters.CharacterFindParameters{
			FindParameters: gFindParameters.FindParameters{
				SliceOptions: gFindParameters.SliceParameters{
					Offset: uint32(c.GetUint("offset")),
					Limit:  uint32(c.GetUint("limit"))}}})

	c.JSON(http.StatusOK,
		result)
}
