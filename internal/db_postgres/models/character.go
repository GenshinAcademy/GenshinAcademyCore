package db_models

// DbCharacters represents Genshin character in database.
type DbCharacter struct {
	Id              DBKey `gorm:"primaryKey"`
	NameId          DBKey `gorm:"column:name"`
	Name            DbString
	CharacterId     GenshinKey `gorm:"column:character_id;uniqueIndex"`
	FullNameId      DBKey      `gorm:"column:full_name"`
	FullName        DbString
	DescriptionId   DBKey `gorm:"column:description"`
	Description     DbString
	TitleId         DBKey `gorm:"column:title"`
	Title           DbString
	Element         uint8              `gorm:"column:element"`
	Rarity          uint8              `gorm:"column:rarity"`
	Weapon          uint8              `gorm:"column:weapon"`
	Icons           []DbCharacterIcon  `gorm:"foreignKey:CharacterId"`
	ArtifactProfits []DbArtifactProfit `gorm:"foreignKey:CharacterId"`
}
