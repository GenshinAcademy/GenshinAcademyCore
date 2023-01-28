package models

import "ga/pkg/core/value_objects/localized_string"

// Language
type Language struct {
	BaseModel
	LanguageName string
}

func (lang Language) CreateNewString(value string) localized_string.LocalizedString {
	if lang.Id == 0 {
		panic("Cannot create new string for not initialized language")
	}

	return localized_string.New(localized_string.LanguageId(lang.Id), value)
}

// Create new string for specified language
func (lang Language) CreateString(src localized_string.LocalizedString, value string) localized_string.LocalizedString {
	if lang.Id == 0 {
		panic("Cannot create new string for not initialized language")
	}

	if src.GetId() == 0 {
		return localized_string.New(localized_string.LanguageId(lang.Id), value)
	}

	return localized_string.Create(src.GetId(), localized_string.LanguageId(lang.Id), value)
}
