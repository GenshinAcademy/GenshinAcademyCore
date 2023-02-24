package models

type ModelId string

const (
	DEFAULT_ID ModelId = ""
)

type BaseModel struct {
	Id ModelId
}

func (id ModelId) IsValid() bool {
	return id != DEFAULT_ID
}
