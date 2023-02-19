package db_models

type DbCharacterIcon struct {
	Id          DBKey `gorm:"primaryKey"`
	CharacterId DBKey
	IconType    byte
	Url         string
}
