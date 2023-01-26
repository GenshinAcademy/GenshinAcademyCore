package models

import "ga/pkg/core/value_objects/localized_string"

// Language
type Language struct {
	BaseModel
	LanguageName string
}

// Create new string for specified language
func (lang Language) CreateNewString(value string) localized_string.LocalizedString {
	if lang.Id == 0 {
		panic("Cannot create new string for not initialized language")
	}
	return localized_string.New(localized_string.LanguageId(lang.Id), value)
}
