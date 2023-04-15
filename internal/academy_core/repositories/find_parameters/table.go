package find_parameters

import (
	"ga/pkg/genshin_core/repositories/find_parameters"
)

type TableSortParameters struct {
	IdSort SortMode
}

type TableFindParameters struct {
	FindParameters
	SliceOptions find_parameters.SliceParameters
	SortParameters TableSortParameters
}

func (parameters TableFindParameters) Slice(offset uint32, limit uint32) TableFindParameters {
	parameters.SliceOptions = find_parameters.SliceParameters{
		Offset: offset,
		Limit:  limit,
	}

	return parameters
}
