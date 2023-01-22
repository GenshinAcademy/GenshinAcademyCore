package db_models

type Db_String struct {
	Id         DBKey `gorm:"primaryKey"`
	LanguageId DBKey
	Value      string
}
