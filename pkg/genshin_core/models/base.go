package models

type ModelId uint64

const (
	UNDEFINED_ID ModelId = 0
)

type BaseModel struct {
	Id ModelId
}

func (id ModelId) IsValid() bool {
	return id != UNDEFINED_ID
}
