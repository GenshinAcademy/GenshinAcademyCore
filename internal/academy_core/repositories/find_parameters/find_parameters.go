package find_parameters

import "ga/internal/academy_core/models"

type SortMode uint8

const(
	SortNone SortMode = iota
	SortByAscending
	SortByDescending
)

type FindParameters struct {
	Ids []models.AcademyId
}

func (param FindParameters) AddId(id models.AcademyId) FindParameters {
	param.Ids = append(param.Ids, id)
	return param
}
