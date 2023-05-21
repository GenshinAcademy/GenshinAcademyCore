package enums

type Gender uint8
type WeaponType uint8
type BodyType uint8
type Element uint8
type Rarity uint8
type Region uint8

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
    // Common 1 star rarity
    Common Rarity = iota
    // Uncommon 2 star rarity
    Uncommon
    // Rare 3 star rarity
    Rare
    // Epic 4 star rarity
	Epic
    // Legendary 5 star rarity
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
