package models

import (
	"ga/pkg/genshin_core/models/languages"
	"ga/pkg/genshindb_wrapper/enums"
)

type CharacterLocalized struct {
	Name    map[languages.Language]string `json:"name" ga:"required,localized"`
	Rarity  enums.Rarity                  `json:"rarity" ga:"required"`
	Element enums.ElementText             `json:"element" ga:"required"`
	Weapon  enums.WeaponText              `json:"weaponType" ga:"required"`
}
