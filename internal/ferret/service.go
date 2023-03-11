package ferret

import (
	"ga/internal/academy_core"
	"ga/internal/academy_core/repositories/find_parameters"
	"ga/internal/ferret/web_models"
	"ga/pkg/genshin_core/models/languages"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type FerretService struct {
	core *academy_core.AcademyCore
}

func CreateService(core *academy_core.AcademyCore) *FerretService {
	var result *FerretService = new(FerretService)
	result.core = core
	return result
}

// GetAllCharactersWithProfits returns all characters profit information in specified language
// Requires Accept-Language header in request
func (service *FerretService) GetAllCharactersWithProfits(c *gin.Context) {
	// TODO: GetProvider should return error if provider is not found
	var language = languages.GetLanguage(languages.ConvertStringsToLanguages(strings.Split(c.GetHeader("Accept-Language"), ",")))

	var characterRepo = service.core.GetProvider(language).NewCharacterRepo()
	var result = characterRepo.FindCharacters(find_parameters.CharacterFindParameters{})
	var characters []web_models.FerretCharacter
	for _, char := range result {
		characters = append(characters, service.mapCharacter(char))
	}

	c.JSON(http.StatusOK,
		characters)
}
