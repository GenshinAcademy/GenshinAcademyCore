package db_repositories

import (
	db_mappers "ga/internal/db_postgres/mappers"
	db_models "ga/internal/db_postgres/models"
	"ga/pkg/core/models"
	"ga/pkg/core/repositories"

	"gorm.io/gorm"
)

// *** Character repository ***//
type PostgresCharacterRepository struct {
	language       string
	gormConnection *gorm.DB
}

func addCharacterPreloads(db *gorm.DB) *gorm.DB {
	return db.
		Joins("Icons")
}

func (repo PostgresCharacterRepository) GetLanguage() string {
	return repo.language
}

func (repo PostgresCharacterRepository) FindCharacterById(characterId models.ModelId) models.Character {
	var selectedCharacter db_models.Db_Character
	addCharacterPreloads(repo.gormConnection).
		First(&selectedCharacter, db_models.DBKey(characterId))

	return db_mappers.CharacterfromDbModel(&selectedCharacter)
}

func (repo PostgresCharacterRepository) FindCharacters(parameters repositories.CharacterFindParameters) []models.Character {

	var selectedChacters []db_models.Db_Character = make([]db_models.Db_Character, 0)
	var result []models.Character = make([]models.Character, 0)
	gormConnection := addCharacterPreloads(repo.gormConnection)

	if len(parameters.Ids) > 0 {
		gormConnection.Find(&selectedChacters, parameters.Ids)
	} else {
		if len(parameters.Names) > 0 {
			gormConnection = gormConnection.Where("character_id IN ?", parameters.Names)
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

func (repo PostgresCharacterRepository) AddCharacter(character models.Character) (models.ModelId, error) {
	panic("TODO")
}

func (repo PostgresCharacterRepository) GetCharacterNames(parameters repositories.CharacterFindParameters) []string {
	var characterNames []db_models.Db_Character
	repo.gormConnection.Select([]string{"characterid"}, &characterNames)
	var result []string = make([]string, 0)
	for _, character := range characterNames {
		result = append(result, character.CharacterId)
	}
	return result
}
