package db_models

// Character represents table for Genshin character in database.
type Character struct {
	Id              DBKey `gorm:"primaryKey"`
	NameId          DBKey `gorm:"column:name"`
	Name            String
	CharacterId     GenshinKey `gorm:"column:character_id;uniqueIndex"`
	FullNameId      DBKey      `gorm:"column:full_name"`
	FullName        String
	DescriptionId   DBKey `gorm:"column:description"`
	Description     String
	TitleId         DBKey `gorm:"column:title"`
	Title           String
	Element         uint8            `gorm:"column:element"`
	Rarity          uint8            `gorm:"column:rarity"`
	Weapon          uint8            `gorm:"column:weapon"`
	Icons           []CharacterIcon  `gorm:"foreignKey:CharacterId"`
	ArtifactProfits []ArtifactProfit `gorm:"foreignKey:CharacterId"`
}
