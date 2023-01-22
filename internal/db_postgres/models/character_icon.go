package db_models

type Db_CharacterIcon struct {
	Id          DBKey `gorm:"primaryKey"`
	CharacterId DBKey
	IconType    byte
	Url         string
}
