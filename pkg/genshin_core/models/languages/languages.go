// Package languages provides information about languages
package languages

// Language is type used for languages
type Language string

const (
	English Language = "en"
	Russian Language = "ru"
)

var (
	// AcceptedLanguages list of accepted languages to operate with database
	AcceptedLanguages = map[Language]bool{
		English: true,
		Russian: true,
	}

	// DefaultLanguage is default language to operate with database
	DefaultLanguage Language = English
)

// GetLanguage returns first found language or default language
func GetLanguage(languages []Language) Language {
	for _, lang := range languages {
		if AcceptedLanguages[lang] {
			return lang
		}
	}

	return DefaultLanguage
}

// ConvertStringsToLanguages converts a slice of strings type to a slice of Language type
func ConvertStringsToLanguages(strSlice []string) []Language {
	langSlice := make([]Language, len(strSlice))
	for i, langStr := range strSlice {
		langSlice[i] = Language(langStr)
	}
	return langSlice
}
