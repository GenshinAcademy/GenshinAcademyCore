package enums

type Gender uint8     //@name Gender
type WeaponType uint8 //@name WeaponType
type BodyType uint8   //@name BodyType
type Element uint8    //@name Element

// Rarity types:
// * 0 - Common 1 star rarity
// * 1 - Uncommon 2 star rarity
// * 2 - Rare 3 star rarity
// * 3 - Epic 4 star rarity
// * 4 - Legendary 5 star rarity
type Rarity uint8 //@name Rarity

type Region uint8 //	@name	Region

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
