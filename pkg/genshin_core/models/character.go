package models

import (
	"ga/pkg/genshin_core/models/enums"
	"ga/pkg/genshin_core/value_objects"
)

type Character struct {
	BaseModel
	Name        string
	FullName    string
	Description string
	Title       string
	Element     enums.Element
	Rarity      enums.Rarity
	Weapon      enums.WeaponType
	Icons       []value_objects.CharacterIcon
	//Body          BodyType
	//Constellation string
	//BirthdayMMDD  string
	//Birthday      string
	//Association   string
	//Affiliation   string
	//Region        Region
}
