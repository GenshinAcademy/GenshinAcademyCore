package db_models

// DbString represents table for string refs and other entities in database.
type DbString struct {
	Id           DBKey           `gorm:"primaryKey;autoIncrement:true"`
	StringValues []DbStringvalue `gorm:"foreignKey:Id"`
}

// DbStringvalue represents string values. Separate table in database.
type DbStringvalue struct {
	Id         DBKey `gorm:"primaryKey"`
	LanguageId DBKey `gorm:"primaryKey"`
	Language   DbLanguage
	Value      string
}

// GetValue gets first string value.
func (str DbString) GetValue() string {
	if len(str.StringValues) == 0 {
		return ""
	}
	return str.StringValues[0].Value
}

// GetLanguageId gets first language.
func (str DbString) GetLanguageId() DBKey {
	if len(str.StringValues) == 0 {
		return DBKey(0)
	}
	return str.StringValues[0].LanguageId
}
