package db_models

// Language represents the language for strings in database.
type Language struct {
	Id DBKey `gorm:"column:id;primaryKey;"`
	//FlagIcon sql.NullString
	Name string `gorm:"column:name;uniqueIndex"`
}
