package main

import (
	"fmt"
	core "ga/pkg/core"
	"ga/pkg/core/models"
	"ga/pkg/core/repositories"
	"ga/pkg/core/value_objects/localized_string"
	db "ga/pkg/db_postgres"
)

func createCharacter(language models.Language, charRepo *repositories.ICharacterRepository) {
	var languageId localized_string.LanguageId = localized_string.LanguageId(language.Id)
	var hutao = models.Character{
		CharacterId: "hu_tao",
		Name:        localized_string.New(languageId, "Hu Tao"),
		FullName:    localized_string.New(languageId, "Hu Taoooo"),
		Description: localized_string.New(languageId, "Pyro damage dealer"),
		Title:       localized_string.New(languageId, "SomeTitle"),
		Element:     models.Pyro,
		Rarity:      models.Legendary,
		Weapon:      models.Polearm,
	}
	(*charRepo).AddCharacter(hutao)
}

func getPostgresRepo() repositories.IRepositoryProvider {
	return db.CreatePostgresProvider()
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
	var config core.GenshinCoreConfiguration = core.GenshinCoreConfiguration{}

	var gacore *core.GenshinCore = core.CreateGenshinCore(config)
	gacore.SetProviderFunc(core.GetProviderFunc(getPostgresRepo))

	var provider = gacore.GetProvider("English")
	langRepo := provider.NewLanguageRepo()
	var lang = langRepo.FindLanguage("English")

	var charRepo = provider.NewCharacterRepo()

	createCharacter(lang, &charRepo)
	return

	var charNames = charRepo.GetCharacterNames(repositories.CharacterFindParameters{})

	for _, names := range charNames {
		fmt.Println(names)
	}
}
