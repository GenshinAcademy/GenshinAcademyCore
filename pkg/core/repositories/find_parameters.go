package repositories

import "ga/pkg/core/models"

type FindParameters struct {
	Ids []models.ModelId
}

func (param *FindParameters) AddId(id models.ModelId) *FindParameters {
	param.Ids = append(param.Ids, id)
	return param
}

type CharacterFindParameters struct {
	FindParameters
	Names    []string
	Elements []models.Element
}

func (param *CharacterFindParameters) AddName(name string) *CharacterFindParameters {
	param.Names = append(param.Names, name)
	return param
}

func (param *CharacterFindParameters) AddElement(element models.Element) *CharacterFindParameters {
	param.Elements = append(param.Elements, element)
	return param
}
