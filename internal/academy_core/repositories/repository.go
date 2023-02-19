package repositories

import (
	"ga/internal/academy_core/models"
	"ga/internal/academy_core/repositories/find_parameters"
    genshin_models "ga/pkg/genshin_core/models"
)

type IRepository interface {
	GetLanguage() models.Language
}

type ILanguageRepository interface {
	AddLanguage(language *models.Language)
	FindLanguage(lang string) models.Language
}

type ICharacterRepository interface {
	IRepository
	GetCharacterIds(parameters find_parameters.CharacterFindParameters) []string
    FindCharacterById(characterId models.AcademyId) (models.Character, bool)
    FindCharacterByGenshinId(characterId genshin_models.ModelId) (models.Character, bool)
	FindCharacters(parameters find_parameters.CharacterFindParameters) []models.Character
	AddCharacter(character *models.Character) (models.AcademyId, error)
	UpdateCharacter(character *models.Character)
}

/*type ICharacterIconRepository interface {
	FindIconsByCharacterId(characterId models.ModelId) []models.CharacterIcon
}*/

type IRepositoryProvider interface {
	GetLanguage() models.Language
	NewCharacterRepo() ICharacterRepository
	//NewCharacterIconRepo() ICharacterIconRepository
}
