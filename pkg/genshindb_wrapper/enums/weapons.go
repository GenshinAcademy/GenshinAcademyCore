package enums

type WeaponText string

const (
	UndefinedWeapon WeaponText = "None"
	Sword           WeaponText = "Sword"
	Claymore        WeaponText = "Claymore"
	Polearm         WeaponText = "Polearm"
	Bow             WeaponText = "Bow"
	Catalyst        WeaponText = "Catalyst"
)

type WeaponType string

// TODO: implement
const (
	SwordType    WeaponType = "WEAPON_SWORD_ONE_HAND"
	ClaymoreType WeaponType = "WEAPON_CLAYMORE"
	PolearmType  WeaponType = "WEAPON_POLE"
	BowType      WeaponType = "WEAPON_BOW"
	CatalystType WeaponType = "WEAPON_CATALYST"
)
