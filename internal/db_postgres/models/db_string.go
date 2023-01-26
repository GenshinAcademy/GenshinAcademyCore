package db_models

//Special table for connectiong string refs and other entities in DB
type Db_String struct {
	Id           DBKey            `gorm:"primaryKey;autoIncrement:true"`
	StringValues []Db_StringValue `gorm:"foreignKey:Id"`
}

// Value of string. Separate table in DB
type Db_StringValue struct {
	Id         DBKey `gorm:"primaryKey"`
	LanguageId DBKey `gorm:"primaryKey"`
	Language   Db_Language
	Value      string
}

// Gets first string value
func (str Db_String) GetValue() string {
	if len(str.StringValues) == 0 {
		return ""
	}
	return str.StringValues[0].Value
}

// Gets first language
func (str Db_String) GetLanguageId() DBKey {
	if len(str.StringValues) == 0 {
		return DBKey(0)
	}
	return str.StringValues[0].LanguageId
}
