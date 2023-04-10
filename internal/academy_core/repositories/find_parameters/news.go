package find_parameters

import (
	"ga/pkg/genshin_core/repositories/find_parameters"
	"time"
)

type NewsFindParameters struct {
	FindParameters
	SliceOptions         find_parameters.SliceParameters
	PublishTimeFrom      *time.Time
	PublishTimeTo        *time.Time
	SortByDescendingTime bool
}

func (parameters NewsFindParameters) Slice(offset uint32, limit uint32) NewsFindParameters {
	parameters.SliceOptions = find_parameters.SliceParameters{
		Offset: offset,
		Limit:  limit,
	}

	return parameters
}
