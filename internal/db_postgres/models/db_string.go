package db_models

// DbString Special table for connectiong string refs and other entities in DB
type DbString struct {
	Id           DBKey           `gorm:"primaryKey;autoIncrement:true"`
	StringValues []DbStringvalue `gorm:"foreignKey:Id"`
}

// DbStringvalue Value of string. Separate table in DB
type DbStringvalue struct {
	Id         DBKey `gorm:"primaryKey"`
	LanguageId DBKey `gorm:"primaryKey"`
	Language   DbLanguage
	Value      string
}

// GetValue Gets first string value
func (str DbString) GetValue() string {
	if len(str.StringValues) == 0 {
		return ""
	}
	return str.StringValues[0].Value
}

// GetLanguageId Gets first language
func (str DbString) GetLanguageId() DBKey {
	if len(str.StringValues) == 0 {
		return DBKey(0)
	}
	return str.StringValues[0].LanguageId
}
