package find_parameters

import (
    "ga/pkg/genshin_core/models"
    "ga/pkg/genshin_core/models/enums"
)

type CharacterFindParameters struct {
    FindParameters
    Elements []enums.Element
}

func (param CharacterFindParameters) AddElement(element enums.Element) CharacterFindParameters {
	param.Elements = append(param.Elements, element)
	return param
}

func FindByCharacterId(characterId string) CharacterFindParameters {
	return CharacterFindParameters{
		FindParameters: FindParameters{
			Ids: []models.ModelId{models.ModelId(characterId)},
		},
	}
}
