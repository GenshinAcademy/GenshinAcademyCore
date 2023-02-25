package db_models

// DbCharacterIcons represents icons for Genshin characters in database.
type DbCharacterIcon struct {
	CharacterId DBKey `gorm:"primaryKey"`
	IconType    uint8 `gorm:"primaryKey"`
	Url         string
}

// TODO: Make URL type that checks format
