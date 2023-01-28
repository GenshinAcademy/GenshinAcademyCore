package main

import (
	"fmt"
	core "ga/pkg/core"
	"ga/pkg/core/models"
	"ga/pkg/core/repositories"
	"ga/pkg/core/repositories/find_parameters"
	"ga/pkg/core_postgres"
)

// Sample method for creating characters
func createHuTao(language models.Language, charRepo *repositories.ICharacterRepository) {
	//var languageId localized_string.LanguageId = localized_string.LanguageId(language.Id)
	//localized_string.New(languageId, "SomeTitle"), Is valid option to craete strings too!
	var hutao = models.Character{
		CharacterId: "hu_tao",
		Element:     models.Pyro,
		Rarity:      models.Legendary,
		Weapon:      models.Polearm,
	}

	hutao.Name = language.CreateString(hutao.Name, "Hu Tao")
	hutao.FullName = language.CreateString(hutao.FullName, "Hu Taoooo")
	hutao.Description = language.CreateString(hutao.Description, "Pyro character")
	hutao.Title = language.CreateString(hutao.Title, "Title")

	(*charRepo).AddCharacter(&hutao)
}

func updateCharacter(language models.Language, character *models.Character, repo repositories.ICharacterRepository) {
	character.Name = language.CreateString(character.Name, "Name"+language.LanguageName)
	character.FullName = language.CreateString(character.FullName, "FullName"+language.LanguageName)
	character.Description = language.CreateString(character.Description, "Description"+language.LanguageName)
	character.Title = language.CreateString(character.Title, "Title"+language.LanguageName)
	repo.UpdateCharacter(character)
}

func main() {
	var dbConfig core_postgres.PostgresDatabaseConfiguration = core_postgres.PostgresDatabaseConfiguration{
		Host:         "localhost",
		UserName:     "postgres",
		UserPassword: "12345678",
		Name:         "gacore_db",
		Port:         "5432",
		ServerPort:   "",
	}
	core_postgres.InitializePostgresDatabase(dbConfig)
	defer core_postgres.CleanupConnections()

	var defaultLanguage = "Russian"
	//Initializing gacore config and configure it for postgres db
	var config core.GenshinCoreConfiguration = core.GenshinCoreConfiguration{
		DefaultLanguage: defaultLanguage,
	}
	core_postgres.ConfigurePostgresDB(&config)

	//Create ga core
	var gacore *core.GenshinCore = core.CreateGenshinCore(config)

	var langRepo = gacore.GetLanguageRepository()

	//Create language if it does not exist
	var language = langRepo.FindLanguage(gacore.GetDefaultLanguageName())
	var upd bool = false
	if language.Id == 0 {
		language = models.Language{
			LanguageName: defaultLanguage,
		}
		langRepo.AddLanguage(&language)
		fmt.Println("Language created successfully!\n", language)
		upd = true
	} else {
		fmt.Println("Language found soccessfully!\n", language)
	}

	//Get provider (with default language, and then character repo)
	var provider = gacore.GetDefaultProvider()
	var characterRepo = provider.NewCharacterRepo()

	var findParams = find_parameters.FindByCharacterId("hu_tao")
	var result = characterRepo.FindCharacters(findParams)
	var hutao models.Character = characterRepo.FindCharacterById(1)
	if len(result) == 0 {
		createHuTao(language, &characterRepo)
		var hutaoNew = characterRepo.FindCharacters(findParams)[0]
		fmt.Println("Hu tao successfully added to DB!\n", hutaoNew)
		return
	}
	if upd {
		updateCharacter(language, &hutao, characterRepo)
		fmt.Println("Hu Tao model updated successfully!\n", hutao)
	}

	fmt.Println("Hu Tao model retrieved successfully!\n", hutao)
}
