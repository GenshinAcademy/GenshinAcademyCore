package db_models

type Db_Character struct {
	Id            DBKey              `gorm:"primaryKey"`
	NameId        DBKey              `gorm:"column:name"`
	CharacterId   string             `gorm:"column:character_id;uniqueIndex"`
	FullNameId    DBKey              `gorm:"column:full_name"`
	DescriptionId DBKey              `gorm:"column:description"`
	TitleId       DBKey              `gorm:"column:title"`
	Element       byte               `gorm:"column:element"`
	Rarity        byte               `gorm:"column:rarity"`
	Weapon        byte               `gorm:"column:weapon"`
	Icons         []Db_CharacterIcon `gorm:"foreignKey:CharacterId"`
}
