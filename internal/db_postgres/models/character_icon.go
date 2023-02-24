package db_models

type DbCharacterIcon struct {
	CharacterId DBKey `gorm:"primaryKey"`
	IconType    uint8  `gorm:"primaryKey"`
	Url         string
}

// TODO: Make URL type that checks format 