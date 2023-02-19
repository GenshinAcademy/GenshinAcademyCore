package db_models

type DbLanguage struct {
	Id DBKey `gorm:"column:id;primaryKey;"`
	//FlagIcon sql.NullString
	Name string `gorm:"column:name;uniqueIndex"`
}
