package models

type CharacterIcon struct {
	BaseModel
	Type CharacterIconType
	Url  string
}
