package models

import (
	"ga/internal/types"
)

type Character struct {
	Id          types.CharacterId
	Name        string
	Description string
	Rarity      types.Rarity
	Element     types.Element
	WeaponType  types.WeaponType
	IconsUrl    types.IconsUrl
}

type CharacterMultilingual struct {
	Id          types.CharacterId
	Name        types.LocalizedString
	Description types.LocalizedString
	Rarity      types.Rarity
	Element     types.Element
	WeaponType  types.WeaponType
	IconsUrl    types.IconsUrl
}
