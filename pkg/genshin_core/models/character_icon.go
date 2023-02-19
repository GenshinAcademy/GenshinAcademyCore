package models

type CharacterIconType byte

const (
	Front CharacterIconType = iota
	Side
	Gacha
)

type CharacterIcon struct {
	CharacterId ModelId
	Type        CharacterIconType
	Url         string
}
