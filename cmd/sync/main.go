package main

import (
	"ga/internal/academy_core"
	academy_models "ga/internal/academy_core/models"
	"ga/internal/configuration"
	academy_postgres "ga/internal/db_postgres/implementation/academy"
	core "ga/pkg/genshin_core"
	gc_models "ga/pkg/genshin_core/models"
	gc_enums "ga/pkg/genshin_core/models/enums"
	"ga/pkg/genshin_core/models/languages"
	"ga/pkg/genshin_core/value_objects"
	"ga/pkg/genshindb_wrapper"
	gdb_enums "ga/pkg/genshindb_wrapper/enums"
	gdb_models "ga/pkg/genshindb_wrapper/models"
	"strings"
)

func main() {
	configuration.Init()

	logger := configuration.Logger

	defer configuration.Logger.Sync()

	apiEn := genshindb_wrapper.Create("https://genshin-db-api.vercel.app/api", gdb_enums.English, logger.Sugar())
	apiRu := genshindb_wrapper.Create("https://genshin-db-api.vercel.app/api", gdb_enums.Russian, logger.Sugar())

	characters, err := apiEn.GetAllCharacters()

	if err != nil {
		panic(err)
	}

	var dbConfig academy_postgres.PostgresDatabaseConfiguration = academy_postgres.PostgresDatabaseConfiguration{
		Host:         configuration.ENV.DBHost,
		UserName:     configuration.ENV.DBUserName,
		UserPassword: configuration.ENV.DBUserPassword,
		DatabaseName: configuration.ENV.DBName,
		Port:         configuration.ENV.DBPort,
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

	//Create enLanguage if it does not exist
	var enLanguage = gacore.GetDefaultLanguage()

	if enLanguage.Id == 0 {
		enLanguage = academy_models.Language{
			LanguageName: string(defaultLanguage),
		}
		langRepo.AddLanguage(&enLanguage)
		logger.Sugar().Infow("Language created successfully!",
			"language", enLanguage)
	} else {
		logger.Sugar().Infow("Language found successfully!",
			"language", enLanguage)
	}

	var characterRepoEn = gacore.AsGenshinCore().GetDefaultProvider().NewCharacterRepo()

	//Get provider for russian language
	var ruLanguage = gacore.GetLanguageRepository().FindLanguage(languages.Russian)

	if ruLanguage.Id == 0 {
		ruLanguage = academy_models.Language{
			LanguageName: string(languages.Russian),
		}
		langRepo.AddLanguage(&ruLanguage)
		logger.Sugar().Infow("Language created successfully!",
			"language", ruLanguage)
	} else {
		logger.Sugar().Infow("Language found successfully!",
			"language", ruLanguage)
	}

	var characterRepoRu = gacore.AsGenshinCore().GetProvider(languages.Russian).NewCharacterRepo()

	for _, character := range characters {
		char := convertCharacter(character)
		characterRepoEn.AddCharacter(&char)
		ruCharFromRepo, _ := characterRepoRu.FindCharacterById(char.Id)

		charRu, err := apiRu.GetCharacter(char.Name)
		if err != nil {
			panic(err)
		}

		addStrings(charRu, &ruCharFromRepo)

		characterRepoRu.UpdateCharacter(&ruCharFromRepo)
	}
}

// convertCharacter converts character from genshin-db by theBowja to genshin-core model
func convertCharacter(input gdb_models.Character) (output gc_models.Character) {
	output.Id = gc_models.ModelId(strings.ToLower(strings.ReplaceAll(input.Name, " ", "_")))

	addStrings(input, &output)

	switch input.Element {
	case "Geo":
		output.Element = gc_enums.Geo
	case "Dendro":
		output.Element = gc_enums.Dendro
	case "Cryo":
		output.Element = gc_enums.Cryo
	case "Pyro":
		output.Element = gc_enums.Pyro
	case "Hydro":
		output.Element = gc_enums.Hydro
	case "Electro":
		output.Element = gc_enums.Electro
	case "Anemo":
		output.Element = gc_enums.Anemo
	default:
		output.Element = gc_enums.UndefinedElement
	}

	switch input.Rarity {
	case "5":
		output.Rarity = gc_enums.Legendary
	default:
		output.Rarity = gc_enums.Epic
	}

	switch input.Weapontype {
	case "Sword":
		output.Weapon = gc_enums.Sword
	case "Bow":
		output.Weapon = gc_enums.Bow
	case "Claymore":
		output.Weapon = gc_enums.Claymore
	case "Catalyst":
		output.Weapon = gc_enums.Catalyst
	case "Polearm":
		output.Weapon = gc_enums.Polearm
	default:
		output.Weapon = gc_enums.UndefinedWeapon
	}

	output.Icons = []value_objects.CharacterIcon{{Type: 0, Url: input.Images.Icon}}
	return output
}

func addStrings(input gdb_models.Character, output *gc_models.Character) {
	output.Name = input.Name
	output.FullName = input.FullName
	output.Description = input.Description
	output.Title = input.Title
}
