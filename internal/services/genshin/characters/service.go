package characters

import (
	"ga/internal/academy_core"
	"ga/internal/academy_core/repositories"
	"ga/internal/services/genshin/characters/models"
	"ga/internal/services/handlers"
	gc_models "ga/pkg/genshin_core/models"
	gc_enums "ga/pkg/genshin_core/models/enums"
	"ga/pkg/genshin_core/models/languages"
	"ga/pkg/genshin_core/repositories/find_parameters"
	"ga/pkg/genshin_core/value_objects"
	gdb_enums "ga/pkg/genshindb_wrapper/enums"
	gdb_models "ga/pkg/genshindb_wrapper/models"

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

// GetAllCharacters godoc
// @Summary Get all characters from database
// @Tags characters
// @Description Retrieves all characters.
// @Produce json
// @Param Accept-Languages header string true "Result language" default(en)
// @Param offset query int false "Offset for pagination"
// @Param limit query int false "Limit for pagination"
// @Success 200 {array} gc_models.Character
// @Failure 404 {error} error "error"
// @Router /characters [get]
func (service *Service) GetAll(c *gin.Context) {
	// TODO: GetProvider should return error if provider is not found
	var language = languages.GetLanguage(languages.ConvertStringsToLanguages(strings.Split(c.GetHeader("Accept-Languages"), ",")))

	var characterRepo = service.core.AsGenshinCore().GetProvider(language).NewCharacterRepo()
	var result, err = characterRepo.FindCharacters(
		find_parameters.CharacterFindParameters{
			FindParameters: find_parameters.FindParameters{
				SliceOptions: find_parameters.SliceParameters{
					Offset: uint32(c.GetUint("offset")),
					Limit:  uint32(c.GetUint("limit"))}}})

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
	}

	c.JSON(http.StatusOK,
		result)
}

// CreateCharacter godoc
// @Summary Add genshin character to database
// @Tags characters
// @Description Creates a new character.
// @Accept json
// @Produce json
// @Param Accept-Languages header string true "Languages splitted by comma. Specify each language you are adding in json body" default(en,ru)
// @Param character body models.CharacterLocalized true "Character data"
// @Security ApiKeyAuth
// @Router /characters [post]
// @Success 200 {array} gc_models.Character
// @Failure 400 {string} string "error"
// @Failure 500 {object} string "error"
func (service *Service) Create(c *gin.Context) {
	// Get languages repositories
	langs := languages.ConvertStringsToLanguages(strings.Split(c.GetHeader("Accept-Languages"), ","))
	charactersRepos := make(map[languages.Language]repositories.ICharacterRepository, len(langs))
	for _, lang := range langs {
		if lang == languages.DefaultLanguage {
			continue
		}

		// TODO: GetProvider should return error if provider is not found
		repo := service.core.GetProvider(&lang).NewCharacterRepo()
		charactersRepos[lang] = repo
	}

	// Read request body
	var requestData models.CharacterLocalized
	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request data"})
		return
	}

	if err := handlers.HasAllFields(requestData); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := handlers.HasLocalizedDefault(requestData, languages.DefaultLanguage); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// Create general fields using default repository
	defaultRepo := service.core.AsGenshinCore().GetProvider(&languages.DefaultLanguage).NewCharacterRepo()
	var character gdb_models.Character
	character.Name = requestData.Name[languages.DefaultLanguage]
	character.Rarity = requestData.Rarity
	character.ElementText = requestData.Element
	character.WeaponText = requestData.Weapon

	// Add to database
	var results []gc_models.Character
	char := convertCharacter(character)
	result, err := defaultRepo.AddCharacter(char)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add character", "message": err.Error()})
		return
	}

	results = append(results, result)

	// Update localization fields
	if len(requestData.Name) > 0 {
		errChan := make(chan error)
		for lang, repo := range charactersRepos {
			go func(id gc_models.ModelId, data models.CharacterLocalized, repo repositories.ICharacterRepository, lang languages.Language) {

				res, err := updateLocalizationFields(id, data, repo, lang)
				if err != nil {
					errChan <- err
				}
				results = append(results, res)
				errChan <- nil
			}(result.Id, requestData, repo, lang)
		}
		for range charactersRepos {
			if err := <-errChan; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update localization fields", "message": err.Error()})
				return
			}
		}
	}

	c.JSON(http.StatusOK, results)
}

func convertCharacter(input gdb_models.Character) (output gc_models.Character) {
	output.Id = gc_models.ModelId(strings.ToLower(strings.ReplaceAll(input.Name, " ", "_")))

	addStrings(input, &output)

	switch input.ElementText {
	case gdb_enums.Geo:
		output.Element = gc_enums.Geo
	case gdb_enums.Dendro:
		output.Element = gc_enums.Dendro
	case gdb_enums.Cryo:
		output.Element = gc_enums.Cryo
	case gdb_enums.Pyro:
		output.Element = gc_enums.Pyro
	case gdb_enums.Hydro:
		output.Element = gc_enums.Hydro
	case gdb_enums.Electro:
		output.Element = gc_enums.Electro
	case gdb_enums.Anemo:
		output.Element = gc_enums.Anemo
	default:
		output.Element = gc_enums.UndefinedElement
	}

	switch input.Rarity {
	case 5:
		output.Rarity = gc_enums.Legendary
	default:
		output.Rarity = gc_enums.Epic
	}

	switch input.WeaponText {
	case gdb_enums.Sword:
		output.Weapon = gc_enums.Sword
	case gdb_enums.Bow:
		output.Weapon = gc_enums.Bow
	case gdb_enums.Claymore:
		output.Weapon = gc_enums.Claymore
	case gdb_enums.Catalyst:
		output.Weapon = gc_enums.Catalyst
	case gdb_enums.Polearm:
		output.Weapon = gc_enums.Polearm
	default:
		output.Weapon = gc_enums.UndefinedWeapon
	}

	// TODO: Rework this clown line
	output.Icons = []value_objects.CharacterIcon{{Type: 0, Url: "/characters/icons/" + string(output.Id) + ".png"}}
	return output
}

func addStrings(input gdb_models.Character, output *gc_models.Character) {
	output.Name = input.Name
	output.Description = input.Description
	output.Title = input.Title
}

func updateLocalizationFields(id gc_models.ModelId, requestData models.CharacterLocalized, repo repositories.ICharacterRepository, lang languages.Language) (gc_models.Character, error) {
	// TODO: Error handling
	result, err := repo.FindCharacterByGenshinId(id)
	if err != nil {
		return gc_models.Character{}, err
	}

	if value, ok := requestData.Name[lang]; ok {
		result.Name = value
	}

	newResult, err := repo.UpdateCharacter(result)
	if err != nil {
		return gc_models.Character{}, err
	}

	return newResult.Character, nil
}
