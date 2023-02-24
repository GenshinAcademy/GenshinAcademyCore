package value_objects

type CharacterIconType byte

const (
	Front CharacterIconType = iota
	Side
	Gacha
)

type CharacterIcon struct {
	Type        CharacterIconType
	Url         string
}
