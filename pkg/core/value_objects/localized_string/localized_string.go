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

// Gets string value
func (str LocalizedString) GetValue() string {
	return str.value
}

// Creates emopty string with passed id
func Empty(id StringId) LocalizedString {
	return LocalizedString{
		id:    id,
		value: "",
	}
}

// Creates brand new localized string (Should be used when adding new strings)
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

// Instantiates LocalizedString assuming existing values are passed
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

// Compares string by Id. (Strings of different languages but with same Ids are treated as same strings)
func Equals(str1 *LocalizedString, str2 *LocalizedString) bool {
	if str1 == nil && str2 == nil {
		return true
	}
	if (str1 == nil && str2 != nil) || (str1 != nil && str2 == nil) {
		return false
	}

	return str1.id == str2.id
}

// Compares two strings by Id and language
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

// Works like ToString
func (str LocalizedString) String() string {
	return str.value
}
