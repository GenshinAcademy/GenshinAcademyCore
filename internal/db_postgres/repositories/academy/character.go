package academy

import (
	models "ga/internal/academy_core/models"
	db_mappers "ga/internal/db_postgres/mappers"
	db_models "ga/internal/db_postgres/models"
	"ga/internal/academy_core/repositories/find_parameters"
    genshin_models "ga/pkg/genshin_core/models"

    "gorm.io/gorm"
)

// PostgresCharacterRepository Character repository
type PostgresCharacterRepository struct {
	mapper         db_mappers.Mapper
	language       models.Language
	gormConnection *gorm.DB
}

func (repo PostgresCharacterRepository) preloadStrings(preloads []string) *gorm.DB {
	var connection = repo.gormConnection
	for _, preload := range preloads {
		connection = connection.Preload(preload, "language_id = ?", repo.language.Id)
	}

	return connection
}

var (
	characterStringPreloads []string = []string{
		"Name.StringValues",
		"FullName.StringValues",
		"Description.StringValues",
		"Title.StringValues",
		"Icons",
		"ArtifactProfits",
	}
)

// Automatically adds all preloads
func (repo PostgresCharacterRepository) addCharacterPreloads() *gorm.DB {

	return repo.preloadStrings(characterStringPreloads).
		Preload("Icons")
}

func (repo PostgresCharacterRepository) GetLanguage() models.Language {
	return repo.language
}

func (repo PostgresCharacterRepository) FindCharacterById(characterId models.AcademyId) (models.Character, bool) {
	var selectedCharacter db_models.DbCharacter
	repo.addCharacterPreloads().Where("id = ?", characterId).First(&selectedCharacter)

    return repo.mapper.MapAcademyCharacterFromDbModel(&selectedCharacter), selectedCharacter.Id != models.UNDEFINED_ID
}

func (repo PostgresCharacterRepository) FindCharacterByGenshinId(characterId genshin_models.ModelId) (models.Character, bool) {
    var selectedCharacter db_models.DbCharacter
    repo.addCharacterPreloads().Where("character_id = ?", characterId).First(&selectedCharacter)

    return repo.mapper.MapAcademyCharacterFromDbModel(&selectedCharacter), selectedCharacter.Id != models.UNDEFINED_ID
}

func (repo PostgresCharacterRepository) FindCharacters(parameters find_parameters.CharacterFindParameters) []models.Character {

	var selectedChacters []db_models.DbCharacter = make([]db_models.DbCharacter, 0)
	var result []models.Character = make([]models.Character, 0)
	gormConnection := repo.addCharacterPreloads()

	if len(parameters.Ids) > 0 {
		gormConnection.Find(&selectedChacters, parameters.Ids)
	} else {
		if len(parameters.CharacterFindParameters.Ids) > 0 {
			gormConnection = gormConnection.Where("character_id IN ?", parameters.CharacterFindParameters.Ids)
		}

		if len(parameters.CharacterFindParameters.Elements) > 0 {
			bytes := make([]uint8, 0)
			for _, cByte := range parameters.Elements {
				bytes = append(bytes, uint8(cByte))
			}
			gormConnection = gormConnection.Where("element IN ?", bytes)
		}

		gormConnection.Find(&selectedChacters)
	}

	for _, character := range selectedChacters {
		result = append(result, repo.mapper.MapAcademyCharacterFromDbModel(&character))
	}

	return result
}

func (repo PostgresCharacterRepository) AddCharacter(character *models.Character) (models.AcademyId, error) {
	var newCharacter = repo.mapper.MapDbCharacterFromModel(character)
	repo.gormConnection.Create(&newCharacter)

    return models.AcademyId(newCharacter.Id), nil
}

func (repo PostgresCharacterRepository) UpdateCharacter(character *models.Character) {
	if character.Id == models.UNDEFINED_ID {
		panic("Cannot update not existing character!")
	}

	var characterToUpdate db_models.DbCharacter = repo.mapper.MapDbCharacterFromModel(character)
	repo.gormConnection.Save(&characterToUpdate)
}

func (repo PostgresCharacterRepository) GetCharacterIds(parameters find_parameters.CharacterFindParameters) []genshin_models.ModelId {
	var characterNames []db_models.DbCharacter
	repo.gormConnection.Select([]string{"character_id"}, &characterNames)
    var result []genshin_models.ModelId = make([]genshin_models.ModelId, 0)
	for _, character := range characterNames {
		result = append(result, genshin_models.ModelId(character.CharacterId))
	}
	return result
}
