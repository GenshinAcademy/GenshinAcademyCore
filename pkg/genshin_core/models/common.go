package models

type ModelId string //@name ModelId

const (
	DEFAULT_ID ModelId = ""
)

type BaseModel struct {
	Id ModelId `extensions:"x-order=0"`
} //@name BaseModel

func (id ModelId) IsValid() bool {
	return id != DEFAULT_ID
}
