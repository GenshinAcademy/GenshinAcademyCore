package models

import (
	"ga/pkg/genshin_core/models/enums"
	"ga/pkg/genshin_core/value_objects"
)

type Weapon struct {
	BaseModel
	Name              string
	Description       string
	DescriptionRaw    string
	Rarity            enums.Rarity
	WeaponType        enums.WeaponType
	BaseAttackValue   float64
	MainStatType      string //TODO: Redo to enum?
	MainStatName      string
	BaseStatText      string
	EffectName        string
	EffectTemplateRaw string
	Icons             []value_objects.WeaponIcon
}
