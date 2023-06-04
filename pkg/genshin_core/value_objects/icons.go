package value_objects

type CharacterIconType byte //@name CharacterIconType
type WeaponIconType byte    //@name WeaponIconType

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
	Type CharacterIconType `example:"0"`
	Url  string            `example:"/characters/icons/lisa.webp"`
} //@name CharacterIcon

// WeaponIcon represents icon for weapon
type WeaponIcon struct {
	Type WeaponIconType
	Url  string
} //@name WeaponIcon
