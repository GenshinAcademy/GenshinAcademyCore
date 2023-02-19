package models

type Gender byte
type WeaponType byte
type BodyType byte
type Element byte
type Rarity byte
type Region byte

const (
	Male Gender = iota
	Female
)

const (
	UndefinedWeapon WeaponType = iota
	Sword
	Claymore
	Polearm
	Bow
	Catalyst
)

const (
	UndefinedElement Element = iota
	Pyro
	Hydro
	Geo
	Anemo
	Electro
	Cryo
	Dendro
)

const (
	Common Rarity = iota
	Uncommon
	Rare
	Epic
	Legendary
)

const (
	ChildGirl BodyType = iota
	ChildBoy
	TeenageGirl
	TeenageBoy
	MuscularMan
	Man
	Lady
)

const (
	UnknownRegion Region = iota
	Monstadt
	Liyue
	Inazuma
	Sumeru
	Fontaine
	Natlan
	Snezhnaya
)
