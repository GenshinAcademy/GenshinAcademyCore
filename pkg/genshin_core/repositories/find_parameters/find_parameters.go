package find_parameters

import "ga/pkg/genshin_core/models"

type SliceParameters struct {
    Offset uint32
    Amount uint32
}

func SelectSingle() SliceParameters {
    return SliceParameters{
        Offset: 0,
        Amount: 1,
    }
}

type FindParameters struct {
	Ids []models.ModelId
    SliceOptions SliceParameters
}

func (param FindParameters) AddId(id models.ModelId) FindParameters {
	param.Ids = append(param.Ids, id)
	return param
}
