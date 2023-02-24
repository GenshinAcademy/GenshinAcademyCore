package main

import (
	"fmt"
	"ga/internal/academy_core"
	academy_models "ga/internal/academy_core/models"
	academy_postgres "ga/internal/db_postgres/implementation/academy"
	core "ga/pkg/genshin_core"

	"ga/pkg/genshin_core/models"
	"ga/pkg/genshin_core/models/enums"
	"ga/pkg/genshin_core/models/languages"
	"ga/pkg/genshin_core/repositories"
	"ga/pkg/genshin_core/repositories/find_parameters"
)

// Sample method for creating characters
func createHuTao(charRepo *repositories.CharacterRepository) {
	//var languageId localized_string.LanguageId = localized_string.LanguageId(language.Id)
	//localized_string.New(languageId, "SomeTitle"), Is valid option to craete strings too!
	var hutao = models.Character{
		BaseModel: models.BaseModel{
			Id: models.ModelId("hu_tao"),
		},
		Name:        "Hu Tao",
		FullName:    "Hu Tao",
		Description: "Pyro DPS character",
		Title:       "Some title",
		Element:     enums.Pyro,
		Rarity:      enums.Legendary,
		Weapon:      enums.Polearm,
	}

	(*charRepo).AddCharacter(&hutao)
}

func updateCharacter(character *models.Character, repo repositories.CharacterRepository) {

	character.Name = "Ху Тао"
	character.FullName = "Ху Тао"
	character.Description = "Персонаж тест"
	character.Title = "Титул"

	repo.UpdateCharacter(character)
}

func main() {
	var dbConfig academy_postgres.PostgresDatabaseConfiguration = academy_postgres.PostgresDatabaseConfiguration{
		Host:         "localhost",
		UserName:     "postgres",
		UserPassword: "12345678",
		Name:         "gacore_db",
		Port:         5432,
		ServerPort:   0,
	}

	academy_postgres.InitializePostgresDatabase(dbConfig)
	defer academy_postgres.CleanupConnections()

	var defaultLanguage = languages.English
	//Initializing gacore config and configure it for postgres db
	var config academy_core.AcademyCoreConfiguration = academy_core.AcademyCoreConfiguration{
		GenshinCoreConfiguration: core.GenshinCoreConfiguration{
			DefaultLanguage: defaultLanguage,
		},
	}
	academy_postgres.ConfigurePostgresDB(&config)

	//Create ga core
	var gacore *academy_core.AcademyCore = academy_core.CreateAcademyCore(config)

	var langRepo = gacore.GetLanguageRepository()

	//Create language if it does not exist
	var language = gacore.GetDefaultLanguage()
	var upd bool = false
	if language.Id == 0 {
		language = academy_models.Language{
			LanguageName: string(defaultLanguage),
		}
		langRepo.AddLanguage(&language)
		fmt.Println("Language created successfully!\n", language)
		upd = true
	} else {
		fmt.Println("Language found soccessfully!\n", language)
	}

	//Get provider (with default language, and then character repo)
	var characterRepo = gacore.AsGenshinCore().GetDefaultProvider().NewCharacterRepo()

	var findParams = find_parameters.FindByCharacterId("hu_tao")
	var result = characterRepo.FindCharacters(findParams)
	var hutao, _ = characterRepo.FindCharacterById("hu_tao")
	if len(result) == 0 {
		createHuTao(&characterRepo)
		var hutaoNew = characterRepo.FindCharacters(findParams)[0]
		fmt.Println("Hu tao successfully added to DB!\n", hutaoNew)
		return
	}
	if upd {
		updateCharacter(&hutao, characterRepo)
		fmt.Println("Hu Tao model updated successfully!\n", hutao)
	}

	fmt.Println("Hu Tao model retrieved successfully!\n", hutao)

	var hutaoacad, _ = gacore.GetDefaultProvider().NewCharacterRepo().FindCharacterByGenshinId("hu_tao")
	fmt.Println("Hu Tao ACADEMY model retrieved successfully!\n", hutaoacad)
}
