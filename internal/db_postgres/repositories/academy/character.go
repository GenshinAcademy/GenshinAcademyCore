package academy

import (
	"errors"
	academy_models "ga/internal/academy_core/models"
	"ga/internal/academy_core/repositories/find_parameters"
	"ga/internal/db_postgres"
	"ga/internal/db_postgres/cache"
	db_mappers "ga/internal/db_postgres/mappers"
	db_models "ga/internal/db_postgres/models"
	"ga/internal/db_postgres/repositories"
	genshin_models "ga/pkg/genshin_core/models"

	"gorm.io/gorm"
)

var (
	characterStringPreloads []string = []string{
		"Name.StringValues",
		"FullName.StringValues",
		"Description.StringValues",
		"Title.StringValues",
	}
	characterPreloads []string = []string{
		"Icons",
		"ArtifactProfits",
	}
)

// PostgresCharacterRepository Character repository
type PostgresCharacterRepository struct {
	PostgresBaseRepository
}

func CreatePostgresCharacterRepository(connection *gorm.DB, language *academy_models.Language, cache *cache.Cache) PostgresCharacterRepository {
	return PostgresCharacterRepository{
		PostgresBaseRepository: PostgresBaseRepository{
			language:       language,
			gormConnection: connection,
			mapper:         db_mappers.CreateMapper(language.LanguageName, language, cache),
		},
	}
}

func (repo PostgresCharacterRepository) GetIdField() string {
	return genericIdField
}

func (repo PostgresCharacterRepository) GetPreloads() []string {
	return characterPreloads
}

func (repo PostgresCharacterRepository) GetStringPreloads() []string {
	return characterStringPreloads
}

func (repo PostgresCharacterRepository) FindCharacterById(characterId academy_models.AcademyId) (*academy_models.Character, bool) {
	var selectedCharacter *db_models.Character

	var connection = repositories.CreateQueryBuilder(repo.GetConnection()).
		PreloadAll(repo).
		FilterById(repo, []academy_models.AcademyId{characterId}).
		GetConnection()
	connection.First(selectedCharacter)

	return repo.mapper.MapAcademyCharacterFromDbModel(selectedCharacter), selectedCharacter.Id != db_models.DBKey(academy_models.UNDEFINED_ID)
}

func (repo PostgresCharacterRepository) FindCharacterByGenshinId(characterId genshin_models.ModelId) (*academy_models.Character, bool) {
	var selectedCharacter *db_models.Character

	var connection = repositories.CreateQueryBuilder(repo.GetConnection()).
		PreloadAll(repo).
		GetConnection().
		Where("character_id = ?", characterId)
	connection.First(selectedCharacter)

	return repo.mapper.MapAcademyCharacterFromDbModel(selectedCharacter), selectedCharacter.Id != db_models.DBKey(academy_models.UNDEFINED_ID)
}

func (repo PostgresCharacterRepository) FindCharacters(parameters find_parameters.CharacterFindParameters) []academy_models.Character {

	var selectedChacters = make([]db_models.Character, 0)
	var result = make([]academy_models.Character, 0)
	var queryBuilder = repositories.CreateQueryBuilder(repo.GetConnection()).PreloadAll(repo)

	if len(parameters.Ids) > 0 {
		queryBuilder = queryBuilder.FilterById(repo, parameters.Ids)
	} else {
		var connection = queryBuilder.GetConnection()
		if len(parameters.CharacterFindParameters.Ids) > 0 {
			connection = connection.Where("character_id IN ?", parameters.CharacterFindParameters.Ids)
		}

		if len(parameters.CharacterFindParameters.Elements) > 0 {
			bytes := make([]uint8, 0)
			for _, cByte := range parameters.Elements {
				bytes = append(bytes, uint8(cByte))
			}
			connection = connection.Where("element IN ?", bytes)
		}
		queryBuilder = repositories.CreateQueryBuilder(connection).Slice(&parameters.SliceOptions)
	}

	queryBuilder.GetConnection().Find(&selectedChacters)

	for _, character := range selectedChacters {
		result = append(result, *repo.mapper.MapAcademyCharacterFromDbModel(&character))
	}

	return result
}

func (repo PostgresCharacterRepository) AddCharacter(character *academy_models.Character) (*academy_models.Character, error) {
	if character == nil {
		return nil, errors.New("null value provided")
	}

	var newCharacter = repo.mapper.MapDbCharacterFromModel(character)

	var connection = repositories.CreateQueryBuilder(repo.GetConnection()).
		PreloadAll(repo).
		GetConnection()

	if err := connection.Create(newCharacter).Error; err != nil {
		return nil, err
	}

	db_postgres.GetCache().UpdateCharacterStrings(newCharacter)

	return character, nil
}

func (repo PostgresCharacterRepository) UpdateCharacter(character *academy_models.Character) (*academy_models.Character, error) {
	if character == nil {
		return nil, errors.New("null value provided")
	}

	if character.Id == academy_models.UNDEFINED_ID {
		return nil, errors.New("not existing character provided")
	}

	var characterToUpdate = repo.mapper.MapDbCharacterFromModel(character)

	var connection = repositories.CreateUpdateQueryBuilder(repo.GetConnection()).
		PreloadAll(repo).
		GetConnection()

	if err := connection.Save(characterToUpdate).Error; err != nil {
		return nil, err
	}

	db_postgres.GetCache().UpdateCharacterStrings(characterToUpdate)

	return character, nil
}

func (repo PostgresCharacterRepository) GetCharacterIds(parameters find_parameters.CharacterFindParameters) []genshin_models.ModelId {
	var characterNames []db_models.Character
	repo.gormConnection.Select([]string{"character_id"}, &characterNames)
	var result = make([]genshin_models.ModelId, 0)
	for _, character := range characterNames {
		result = append(result, genshin_models.ModelId(character.CharacterId))
	}
	return result
}
