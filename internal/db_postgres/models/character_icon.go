package db_models

// CharacterIcon represents table for Genshin characters' icons in database.
type CharacterIcon struct {
	CharacterId DBKey `gorm:"primaryKey"`
	IconType    uint8 `gorm:"primaryKey"`
	Url         string `gorm:"type:varchar"`
}

// TODO: Make URL type that checks format
