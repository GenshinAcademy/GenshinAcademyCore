package db_models

//Character DB model
type Db_Character struct {
	Id            DBKey `gorm:"primaryKey"`
	NameId        DBKey `gorm:"column:name"`
	Name          Db_String
	CharacterId   string `gorm:"column:character_id;uniqueIndex"`
	FullNameId    DBKey  `gorm:"column:full_name"`
	FullName      Db_String
	DescriptionId DBKey `gorm:"column:description"`
	Description   Db_String
	TitleId       DBKey `gorm:"column:title"`
	Title         Db_String
	Element       byte               `gorm:"column:element"`
	Rarity        byte               `gorm:"column:rarity"`
	Weapon        byte               `gorm:"column:weapon"`
	Icons         []Db_CharacterIcon `gorm:"foreignKey:CharacterId"`
}
