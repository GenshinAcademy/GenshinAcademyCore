package models

import (
	"ga/pkg/genshin_core/models/enums"
	"ga/pkg/genshin_core/value_objects"
)

type Character struct {
	BaseModel
	Name        string `example:"Lisa" extensions:"x-order=1"`
	FullName    string `example:"Lisa" extensions:"x-order=2"`
	Description string `example:"The languid but knowledgeable Librarian of the Knights of Favonius, deemed by Sumeru Akademiya to be their most distinguished graduate in the past two centuries." extensions:"x-order=3"`
	Title       string `example:"Witch of Purple Rose" extensions:"x-order=4"`

	// Character type:
	// * 0 - Undefined
	// * 1 - Pyro
	// * 2 - Hydro
	// * 3 - Geo
	// * 4 - Anemo
	// * 5 - Electro
	// * 6 - Cryo
	// * 7 - Dendro
	Element enums.Element `example:"5" extensions:"x-order=5"`

	// Rarity types:
	// * 3 - Epic 4 star rarity
	// * 4 - Legendary 5 star rarity
	Rarity enums.Rarity `example:"3" extensions:"x-order=6"`

	// Character weapon type:
	// * 0 - Undefined
	// * 1 - Sword
	// * 2 - Claymore
	// * 3 - Polearm
	// * 4 - Bow
	// * 5 - Catalyst
	Weapon enums.WeaponType `example:"5" extensions:"x-order=7"`

	Icons []value_objects.CharacterIcon
	//Body          BodyType
	//Constellation string
	//BirthdayMMDD  string
	//Birthday      string
	//Association   string
	//Affiliation   string
	//Region        Region
} //@name GenshinCharacter
