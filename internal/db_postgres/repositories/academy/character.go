package academy

import (
	academy_models "ga/internal/academy_core/models"
	"ga/internal/academy_core/repositories/find_parameters"
	"ga/internal/db_postgres"
	"ga/internal/db_postgres/cache"
	db_mappers "ga/internal/db_postgres/mappers"
	db_models "ga/internal/db_postgres/models"
	genshin_models "ga/pkg/genshin_core/models"

	"gorm.io/gorm"
)

// PostgresCharacterRepository Character repository
type PostgresCharacterRepository struct {
	mapper         db_mappers.Mapper
	language       academy_models.Language
	gormConnection *gorm.DB
}

func CreatePostgresCharacterRepository(connection *gorm.DB, language academy_models.Language, cache *cache.Cache) PostgresCharacterRepository {
	return PostgresCharacterRepository{
		gormConnection: connection,
		mapper:         db_mappers.CreateMapper(language.LanguageName, language, cache),
		language:       language,
	}
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
	}
)

// Automatically adds all preloads
func (repo PostgresCharacterRepository) addCharacterPreloads() *gorm.DB {

	return repo.preloadStrings(characterStringPreloads).
		Preload("Icons").
		Preload("ArtifactProfits")
}

func (repo PostgresCharacterRepository) GetLanguage() academy_models.Language {
	return repo.language
}

func (repo PostgresCharacterRepository) FindCharacterById(characterId academy_models.AcademyId) (academy_models.Character, bool) {
	var selectedCharacter db_models.DbCharacter
	repo.addCharacterPreloads().Where("id = ?", characterId).First(&selectedCharacter)

	return repo.mapper.MapAcademyCharacterFromDbModel(&selectedCharacter), selectedCharacter.Id != db_models.DBKey(academy_models.UNDEFINED_ID)
}

func (repo PostgresCharacterRepository) FindCharacterByGenshinId(characterId genshin_models.ModelId) (academy_models.Character, bool) {
	var selectedCharacter db_models.DbCharacter
	repo.addCharacterPreloads().Where("character_id = ?", characterId).First(&selectedCharacter)

	return repo.mapper.MapAcademyCharacterFromDbModel(&selectedCharacter), selectedCharacter.Id != db_models.DBKey(academy_models.UNDEFINED_ID)
}

func (repo PostgresCharacterRepository) FindCharacters(parameters find_parameters.CharacterFindParameters) []academy_models.Character {

	var selectedChacters = make([]db_models.DbCharacter, 0)
	var result = make([]academy_models.Character, 0)
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

func (repo PostgresCharacterRepository) AddCharacter(character *academy_models.Character) (academy_models.AcademyId, error) {
	var newCharacter = repo.mapper.MapDbCharacterFromModel(character)

	repo.gormConnection.Create(&newCharacter)

	db_postgres.GetCache().UpdateCharacterStrings(&newCharacter)

	return academy_models.AcademyId(newCharacter.Id), nil
}

func (repo PostgresCharacterRepository) UpdateCharacter(character *academy_models.Character) {
	if character.Id == academy_models.UNDEFINED_ID {
		panic("Cannot update not existing character!")
	}

	var characterToUpdate = repo.mapper.MapDbCharacterFromModel(character)
	repo.gormConnection.Save(&characterToUpdate)

	db_postgres.GetCache().UpdateCharacterStrings(&characterToUpdate)
}

func (repo PostgresCharacterRepository) GetCharacterIds(parameters find_parameters.CharacterFindParameters) []genshin_models.ModelId {
	var characterNames []db_models.DbCharacter
	repo.gormConnection.Select([]string{"character_id"}, &characterNames)
	var result = make([]genshin_models.ModelId, 0)
	for _, character := range characterNames {
		result = append(result, genshin_models.ModelId(character.CharacterId))
	}
	return result
}
