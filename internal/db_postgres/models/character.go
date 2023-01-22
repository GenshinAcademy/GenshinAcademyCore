package db_models

type Db_Character struct {
	Id            DBKey `gorm:"primaryKey"`
	NameId        DBKey
	CharacterId   string `gorm:"uniqueIndex"`
	FullNameId    DBKey
	DescriptionId DBKey
	TitleId       DBKey
	Element       byte
	Rarity        byte
	Weapon        byte
	Icons         []Db_CharacterIcon
}
