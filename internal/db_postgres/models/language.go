package db_models

type Db_Language struct {
	Id DBKey `gorm:"primaryKey"`
	//FlagIcon sql.NullString
	Name string
}
