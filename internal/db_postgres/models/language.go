package db_models

type Db_Language struct {
	Id DBKey `gorm:"column:id;primaryKey;"`
	//FlagIcon sql.NullString
	Name string `gorm:"column:name;uniqueIndex"`
}
