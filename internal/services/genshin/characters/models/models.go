package models

import (
	"ga/pkg/genshin_core/models/languages"
	"ga/pkg/genshindb_wrapper/enums"
)

type CharacterLocalized struct {
	Name    map[languages.Language]string `json:"name" ga:"required,localized" example:"en:Liza,ru:Лиза"`
	Rarity  enums.Rarity                  `json:"rarity" ga:"required" example:"4"`
	Element enums.ElementText             `json:"element" ga:"required" example:"Electro"`
	Weapon  enums.WeaponText              `json:"weaponType" ga:"required" example:"Catalyst"`
}
