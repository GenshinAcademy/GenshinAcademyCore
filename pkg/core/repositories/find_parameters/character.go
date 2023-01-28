package find_parameters

import "ga/pkg/core/models"

type CharacterFindParameters struct {
	FindParameters
	CharactedIds []string
	Elements     []models.Element
}

func (param CharacterFindParameters) AddCharacterId(name string) CharacterFindParameters {
	param.CharactedIds = append(param.CharactedIds, name)
	return param
}

func (param CharacterFindParameters) AddElement(element models.Element) CharacterFindParameters {
	param.Elements = append(param.Elements, element)
	return param
}

func FindByCharacterId(characterId string) CharacterFindParameters {
	return CharacterFindParameters{
		CharactedIds: []string{characterId},
	}
}
