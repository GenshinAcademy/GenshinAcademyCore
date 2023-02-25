package db_models

// String represents table for string refs and other entities in database.
type String struct {
	Id           DBKey         `gorm:"primaryKey;autoIncrement:true"`
	StringValues []StringValue `gorm:"foreignKey:Id"`
}

// StringValue represents table for string values. Separate table in database.
type StringValue struct {
	Id         DBKey `gorm:"primaryKey"`
	LanguageId DBKey `gorm:"primaryKey"`
	Language   Language
	Value      string
}

// GetValue gets first string value.
func (str String) GetValue() string {
	if len(str.StringValues) == 0 {
		return ""
	}
	return str.StringValues[0].Value
}

// GetLanguageId gets first language.
func (str String) GetLanguageId() DBKey {
	if len(str.StringValues) == 0 {
		return DBKey(0)
	}
	return str.StringValues[0].LanguageId
}
