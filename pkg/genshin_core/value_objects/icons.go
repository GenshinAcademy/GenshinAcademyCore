package value_objects

type CharacterIconType byte
type WeaponIconType byte

const (
	Front CharacterIconType = iota
	Side
	Gacha
)

const (
	WeaponMenuIcon WeaponIconType = iota
)

// CharacterIcon represents icon for Character
type CharacterIcon struct {
	Type CharacterIconType
	Url  string
}

// WeaponIcon represents icon for weapon
type WeaponIcon struct {
	Type WeaponIconType
	Url  string
}
