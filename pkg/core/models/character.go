package models

import (
	"ga/pkg/core/value_objects/localized_string"
)

type Character struct {
	BaseModel
	Name        localized_string.LocalizedString
	CharacterId string
	FullName    localized_string.LocalizedString
	Description localized_string.LocalizedString
	Title       localized_string.LocalizedString
	Element     Element
	Rarity      Rarity
	Weapon      WeaponType
	Icons       []CharacterIcon
	//Body          BodyType
	//Constellation localized_string.LocalizedString
	//BirthdayMMDD  string
	//Birthday      string
	//Association   string
	//Affiliation   string
	//Region        Region
}
