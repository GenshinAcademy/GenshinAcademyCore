package db_models

type Db_String struct {
	Id         DBKey `gorm:"primaryKey;autoIncrement:true"`
	LanguageId DBKey
	Value      string
}
