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
	weaponStringPreloads = []string{
		"Name.StringValues",
		"Description.StringValues",
		"DescriptionRaw.StringValues",
		"EffectName.StringValues",
		"EffectTemplateRaw.StringValues",
	}
	weaponPreloads = []string{
		"Icons",
	}
)

type PostgresWeaponRepository struct {
	PostgresBaseRepository
}

func CreatePostgresWeaponRepository(connection *gorm.DB, language *academy_models.Language, cache *cache.Cache) PostgresWeaponRepository {
	return PostgresWeaponRepository{
		PostgresBaseRepository: PostgresBaseRepository{
			language:       language,
			gormConnection: connection,
			mapper:         db_mappers.CreateMapper(language.LanguageName, language, cache),
		},
	}
}

func (repo PostgresWeaponRepository) GetIdField() string {
	return genericIdField
}

func (repo PostgresWeaponRepository) GetStringPreloads() []string {
	return weaponStringPreloads
}

func (repo PostgresWeaponRepository) GetPreloads() []string {
	return weaponPreloads
}

func (repo PostgresWeaponRepository) GetWeaponIds(parameters find_parameters.WeaponFindParameters) ([]genshin_models.ModelId, error) {
	var weaponNames []db_models.Weapon
	if err := repo.gormConnection.Select([]string{"weapon_id"}, &weaponNames).Error; err != nil {
		return nil, err
	}

	var result = make([]genshin_models.ModelId, 0)
	for _, weapon := range weaponNames {
		result = append(result, genshin_models.ModelId(weapon.WeaponId))
	}

	return result, nil
}

func (repo PostgresWeaponRepository) FindWeaponById(id academy_models.AcademyId) (academy_models.Weapon, error) {
	var weapon db_models.Weapon
	var ids = make([]academy_models.AcademyId, 1)
	ids[0] = id

	var connection = repositories.CreateQueryBuilder(repo.GetConnection()).
		PreloadAll(repo).
		FilterById(repo, ids).
		GetConnection()

	if err := connection.Find(&weapon).Error; err != nil {
		return academy_models.Weapon{}, err
	}

	return repo.mapper.MapAcademyWeaponFromDbModel(weapon), nil
}

func (repo PostgresWeaponRepository) FindWeaponByGenshinId(weaponId genshin_models.ModelId) (academy_models.Weapon, error) {
	var selectedWeapon db_models.Weapon

	var connection = repositories.CreateQueryBuilder(repo.GetConnection()).
		PreloadAll(repo).
		GetConnection().
		Where("weapon_id = ?", weaponId)

	if err := connection.First(&selectedWeapon).Error; err != nil {
		return academy_models.Weapon{}, nil
	}

	return repo.mapper.MapAcademyWeaponFromDbModel(selectedWeapon), nil
}

func (repo PostgresWeaponRepository) FindWeapons(parameters find_parameters.WeaponFindParameters) ([]academy_models.Weapon, error) {
	var selectedWeapons []db_models.Weapon = make([]db_models.Weapon, 0)

	var queryBuilder = repositories.CreateQueryBuilder(repo.GetConnection()).
		PreloadAll(repo)

	if len(parameters.Ids) > 0 {
		queryBuilder = queryBuilder.FilterById(repo, parameters.Ids)
	} else {

		queryBuilder = queryBuilder.Slice(&parameters.SliceOptions)
	}

	if err := queryBuilder.GetConnection().Find(&selectedWeapons).Error; err != nil {
		return nil, err
	}

	var resultWeapons = make([]academy_models.Weapon, len(selectedWeapons))
	for index, weapon := range selectedWeapons {
		resultWeapons[index] = repo.mapper.MapAcademyWeaponFromDbModel(weapon)
	}

	return resultWeapons, nil
}

func (repo PostgresWeaponRepository) AddWeapon(weapon academy_models.Weapon) (academy_models.Weapon, error) {
	var dbWeapon = repo.mapper.MapDbWeaponFromModel(weapon)

	var connection = repositories.CreateQueryBuilder(repo.GetConnection()).
		PreloadAll(repo).
		GetConnection()

	if err := connection.Create(&dbWeapon).Error; err != nil {
		return academy_models.Weapon{}, err
	}
	result := repo.mapper.MapAcademyWeaponFromDbModel(dbWeapon)

	db_postgres.GetCache().UpdateWeaponStrings(dbWeapon)

	return result, nil
}

func (repo PostgresWeaponRepository) UpdateWeapon(weapon academy_models.Weapon) (academy_models.Weapon, error) {
	if weapon.Id == academy_models.UNDEFINED_ID {
		return academy_models.Weapon{}, errors.New("not existing weapon provided")
	}

	var dbWeapon = repo.mapper.MapDbWeaponFromModel(weapon)

	var connection = repositories.CreateUpdateQueryBuilder(repo.GetConnection()).
		PreloadAll(repo).
		GetConnection()

	if err := connection.Save(&dbWeapon).Error; err != nil {
		return academy_models.Weapon{}, err
	}

	db_postgres.GetCache().UpdateWeaponStrings(dbWeapon)

	weapon.Id = academy_models.AcademyId(dbWeapon.Id)

	return weapon, nil
}
