package find_parameters

import "ga/pkg/genshin_core/models"

type FindParameters struct {
	Ids []models.ModelId
}

func (param FindParameters) AddId(id models.ModelId) FindParameters {
	param.Ids = append(param.Ids, id)
	return param
}
