package presenter

import (
	"ga/internal/models"
	"ga/internal/types"
	"strings"
)

func PresentCharacterData(character any) {
	switch character.(type) {
	case *models.CharacterMultilingual:
		c := character.(*models.CharacterMultilingual)
		c.Id = types.CharacterId(strings.ToLower(strings.ReplaceAll(c.Name[types.DefaultLanguage], " ", "_")))
		c.IconsUrl = map[types.IconType]string{
			types.FrontFace: string(c.Id),
		}
	case *models.Character:
		c := character.(*models.Character)
		c.Id = types.CharacterId(strings.ToLower(strings.ReplaceAll(c.Name, " ", "_")))
		c.IconsUrl = map[types.IconType]string{
			types.FrontFace: string(c.Id),
		}
	}
}
