package db_repositories

import (
	db_mappers "ga/internal/db_postgres/mappers"
	db_models "ga/internal/db_postgres/models"
	"ga/pkg/genshin_core/models"
	"ga/pkg/genshin_core/repositories/find_parameters"

	"gorm.io/gorm"
)

// *** Character repository ***//
type PostgresCharacterRepository struct {
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
	characterPreloads []string = []string{
		"Name.StringValues",
		"FullName.StringValues",
		"Description.StringValues",
		"Title.StringValues",
	}
)

// Automatically adds all preloads
func (repo PostgresCharacterRepository) addCharacterPreloads() *gorm.DB {

	return repo.preloadStrings(characterPreloads).
		Preload("Icons")
	/*return repo.gormConnection.

	Preload("Name.StringValues").
	Preload("FullName.StringValues").
	Preload("Description.StringValues").
	Preload("Title.StringValues").
	Preload("Icons")*/
}

func (repo PostgresCharacterRepository) GetLanguage() models.Language {
	return repo.language
}

func (repo PostgresCharacterRepository) FindCharacterById(characterId models.ModelId) models.Character {
	var selectedCharacter db_models.Db_Character
	repo.addCharacterPreloads().First(&selectedCharacter, db_models.DBKey(characterId))

	return db_mappers.CharacterfromDbModel(&selectedCharacter)
}

func (repo PostgresCharacterRepository) FindCharacters(parameters find_parameters.CharacterFindParameters) []models.Character {

	var selectedChacters []db_models.Db_Character = make([]db_models.Db_Character, 0)
	var result []models.Character = make([]models.Character, 0)
	gormConnection := repo.addCharacterPreloads()

	if len(parameters.Ids) > 0 {
		gormConnection.Find(&selectedChacters, parameters.Ids)
	} else {
		if len(parameters.CharactedIds) > 0 {
			gormConnection = gormConnection.Where("character_id IN ?", parameters.CharactedIds)
		}

		if len(parameters.Elements) > 0 {
			bytes := make([]uint8, 0)
			for _, cByte := range parameters.Elements {
				bytes = append(bytes, uint8(cByte))
			}
			gormConnection = gormConnection.Where("element IN ?", bytes)
		}

		gormConnection.Find(&selectedChacters)
	}

	for _, character := range selectedChacters {
		result = append(result, db_mappers.CharacterfromDbModel(&character))
	}

	return result
}

func (repo PostgresCharacterRepository) AddCharacter(character *models.Character) (models.ModelId, error) {
	var newCharacter = db_mappers.DbCharacterFromModel(character)
	repo.gormConnection.Create(&newCharacter)

	return models.ModelId(newCharacter.Id), nil
}

func (repo PostgresCharacterRepository) UpdateCharacter(character *models.Character) {
	if character.Id == models.UNDEFINED_ID {
		panic("Cannot update not existing character!")
	}

	var characterToUpdate db_models.Db_Character = db_mappers.DbCharacterFromModel(character)
	repo.gormConnection.Save(&characterToUpdate)
}

func (repo PostgresCharacterRepository) GetCharacterIds(parameters find_parameters.CharacterFindParameters) []string {
	var characterNames []db_models.Db_Character
	repo.gormConnection.Select([]string{"character_id"}, &characterNames)
	var result []string = make([]string, 0)
	for _, character := range characterNames {
		result = append(result, character.CharacterId)
	}
	return result
}
