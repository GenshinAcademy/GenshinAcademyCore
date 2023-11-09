package genshin_db

type Character struct {
	Name        string     `json:"name"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	WeaponType  WeaponType `json:"weapontype"`
	Rarity      Rarity     `json:"rarity"`
	Element     Element    `json:"element"`
}
