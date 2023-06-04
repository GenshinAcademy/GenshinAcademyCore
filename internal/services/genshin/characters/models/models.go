package models

import (
	"ga/pkg/genshin_core/models/languages"
	"ga/pkg/genshindb_wrapper/enums"
)

type CharacterLocalized struct {
	Name    map[languages.Language]string `json:"name" ga:"required,localized" example:"en:Lisa,ru:Лиза" extensions:"x-order=0"`
	Rarity  enums.Rarity                  `json:"rarity" ga:"required" example:"4" extensions:"x-order=1"`
	Element enums.ElementText             `json:"element" ga:"required" example:"Electro" extensions:"x-order=2"`
	Weapon  enums.WeaponText              `json:"weaponType" ga:"required" example:"Catalyst" extensions:"x-order=3"`
} //@name CharacterLocalized
