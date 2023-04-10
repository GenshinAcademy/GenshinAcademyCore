package find_parameters

import (
	"ga/pkg/genshin_core/repositories/find_parameters"
)

type TableFindParameters struct {
	FindParameters
	SliceOptions find_parameters.SliceParameters
}

func (parameters TableFindParameters) Slice(offset uint32, limit uint32) TableFindParameters {
	parameters.SliceOptions = find_parameters.SliceParameters{
		Offset: offset,
		Limit:  limit,
	}

	return parameters
}
