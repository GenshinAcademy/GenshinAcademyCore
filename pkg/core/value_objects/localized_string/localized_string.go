package localized_string

type StringId uint
type LanguageId uint

type LocalizedString struct {
	id         StringId
	languageId LanguageId
	value      string
}

func (str LocalizedString) IsEmpty() bool {
	return str.value != ""
}

func (str LocalizedString) GetId() StringId {
	return str.id
}

func (str LocalizedString) GetLanguageId() LanguageId {
	return str.languageId
}

func (str LocalizedString) GetValue() string {
	return str.value
}

func Empty(id StringId) LocalizedString {
	return LocalizedString{
		id:    id,
		value: "",
	}
}

func New(languageId LanguageId, value string) LocalizedString {
	if value == "" {
		panic("Cannot create empty string")
	}
	return LocalizedString{
		id:         0,
		languageId: languageId,
		value:      value,
	}
}

func Create(id StringId, languageId LanguageId, value string) LocalizedString {
	if value == "" {
		panic("Cannot create empty string")
	}
	return LocalizedString{
		id:         id,
		languageId: languageId,
		value:      value,
	}
}

func Equals(str1 *LocalizedString, str2 *LocalizedString) bool {
	if str1 == nil && str2 == nil {
		return true
	}
	if (str1 == nil && str2 != nil) || (str1 != nil && str2 == nil) {
		return false
	}

	return str1.id == str2.id
}

func EqualsCulture(str1 *LocalizedString, str2 *LocalizedString) bool {
	if str1 == nil && str2 == nil {
		return true
	}
	if (str1 == nil && str2 != nil) || (str1 != nil && str2 == nil) {
		return false
	}
	if str1.IsEmpty() || str2.IsEmpty() {
		return false
	}

	return str1.id == str2.id && str1.languageId == str2.languageId
}
