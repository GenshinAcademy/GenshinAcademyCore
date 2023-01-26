package main

import (
	"fmt"
	core "ga/pkg/core"
	"ga/pkg/core/models"
	"ga/pkg/core/repositories"
	db "ga/pkg/db_postgres"
)

// Sample method for creating characters
func createHuTao(language models.Language, charRepo *repositories.ICharacterRepository) {
	//var languageId localized_string.LanguageId = localized_string.LanguageId(language.Id)
	//localized_string.New(languageId, "SomeTitle"), Is valid option to craete strings too!
	var hutao = models.Character{
		CharacterId: "hu_tao",
		Name:        language.CreateNewString("Hu Tao"),
		FullName:    language.CreateNewString("Hu Taoooo"),
		Description: language.CreateNewString("Pyro damage dealer"),
		Title:       language.CreateNewString("SomeTitle"),
		Element:     models.Pyro,
		Rarity:      models.Legendary,
		Weapon:      models.Polearm,
	}
	(*charRepo).AddCharacter(&hutao)
}

func main() {
	var dbConfig db.PostgresDatabaseConfiguration = db.PostgresDatabaseConfiguration{
		Host:         "localhost",
		UserName:     "postgres",
		UserPassword: "12345678",
		Name:         "gacore_db",
		Port:         "5432",
		ServerPort:   "",
	}
	db.InitializePostgresDatabase(dbConfig)
	defer db.CleanupConnections()

	var defaultLanguage = "English"
	//Initializing gacore config and configure it for postgres db
	var config core.GenshinCoreConfiguration = core.GenshinCoreConfiguration{
		DefaultLanguage: defaultLanguage,
	}
	db.ConfigurePostgresDB(&config)

	//Create ga core
	var gacore *core.GenshinCore = core.CreateGenshinCore(config)

	var langRepo = gacore.GetLanguageRepository()

	//Create language if it does not exist
	var language = langRepo.FindLanguage(gacore.GetDefaultLanguageName())
	if language.Id == 0 {
		language = models.Language{
			LanguageName: defaultLanguage,
		}
		langRepo.AddLanguage(&language)
		fmt.Println("Language created successfully!\n", language)
	} else {
		fmt.Println("Language found soccessfully!\n", language)
	}

	//Get provider (with default language, and then character repo)
	var provider = gacore.GetDefaultProvider()
	var characterRepo = provider.NewCharacterRepo()

	var hutao = characterRepo.FindCharacterById(1)
	if hutao.Id == 0 {
		createHuTao(language, &characterRepo)
		var hutaoNew = characterRepo.FindCharacterById(1)
		fmt.Println("Hu tao successfully added to DB!\n", hutaoNew)
		return
	}

	fmt.Println("Hu Tao model retrieved successfully!\n", hutao)
}
