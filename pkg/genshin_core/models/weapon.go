package models

import "ga/pkg/genshin_core/models/enums"

type Weapon struct {
	BaseModel
	Name            string
	Description     string
	DescriptionRaw  string
	Rarity          enums.Rarity
	WeaponType      enums.WeaponType
	BaseAttackValue float64
	WeaponStatType  string //TODO: Redo to enum?
	EffectName      string
}
