package models

type Character struct {
	BaseModel
	Name        string
	CharacterId string
	FullName    string
	Description string
	Title       string
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
