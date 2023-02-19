package find_parameters

import "ga/pkg/genshin_core/models"

type CharacterFindParameters struct {
	FindParameters
	Elements []models.Element
}

func (param CharacterFindParameters) AddElement(element models.Element) CharacterFindParameters {
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
