package genshin

import (
	academy_find_parameters "ga/internal/academy_core/repositories/find_parameters"
	"ga/pkg/genshin_core/errors"

	genshin_models "ga/pkg/genshin_core/models"
	"ga/pkg/genshin_core/models/languages"
	find_parameters "ga/pkg/genshin_core/repositories/find_parameters"

	"ga/internal/db_postgres/repositories/academy"
)

type PostgresGenshinWeaponRepository struct {
	academyRepo *academy.PostgresWeaponRepository
}

func CreateGenshinWeaponRepository(repo *academy.PostgresWeaponRepository) *PostgresGenshinWeaponRepository {
	var genshinRepo = new(PostgresGenshinWeaponRepository)
	genshinRepo.academyRepo = repo

	return genshinRepo
}
func (repo PostgresGenshinWeaponRepository) GetLanguage() *languages.Language {
	result := languages.Language(repo.academyRepo.GetLanguage().LanguageName)
	return &result
}

func (repo PostgresGenshinWeaponRepository) GetWeaponIds(parameters find_parameters.WeaponFindParameters) ([]genshin_models.ModelId, error) {
	var academyParams = academy_find_parameters.WeaponFindParameters{
		WeaponFindParameters: parameters,
	}

	return repo.academyRepo.GetWeaponIds(academyParams)
}

func (repo PostgresGenshinWeaponRepository) FindWeaponById(weaponId genshin_models.ModelId) (genshin_models.Weapon, error) {
	var weapon, err = repo.academyRepo.FindWeaponByGenshinId(weaponId)
	if err != nil {
		return genshin_models.Weapon{}, errors.WeaponNotFound(weaponId)
	}

	return weapon.Weapon, nil
}

func (repo PostgresGenshinWeaponRepository) FindWeapons(parameters find_parameters.WeaponFindParameters) ([]genshin_models.Weapon, error) {
	var weapons = make([]genshin_models.Weapon, 0)
	var academyParams = academy_find_parameters.WeaponFindParameters{
		WeaponFindParameters: parameters,
	}

	var academyWeapons, err = repo.academyRepo.FindWeapons(academyParams)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(academyWeapons); i += 1 {
		weapons = append(weapons, academyWeapons[i].Weapon)
	}

	return weapons, nil
}

func (repo PostgresGenshinWeaponRepository) AddWeapon(model genshin_models.Weapon) (genshin_models.Weapon, error) {
	if model.Id == genshin_models.DEFAULT_ID {
		return genshin_models.Weapon{}, errors.EmptyId()
	}

	var academyWeapon, _ = repo.academyRepo.FindWeaponByGenshinId(model.Id)
	academyWeapon.Weapon = model
	result, err := repo.academyRepo.AddWeapon(academyWeapon)
	if err != nil {
		return genshin_models.Weapon{}, err
	}

	return result.Weapon, nil
}

func (repo PostgresGenshinWeaponRepository) UpdateWeapon(model genshin_models.Weapon) (genshin_models.Weapon, error) {
	if model.Id == genshin_models.DEFAULT_ID {
		return genshin_models.Weapon{}, errors.EmptyId()
	}

	var academyWeapon, err = repo.academyRepo.FindWeaponByGenshinId(model.Id)
	if err != nil {
		return genshin_models.Weapon{}, errors.WeaponNotFound(model.Id)
	}

	academyWeapon.Weapon = model
	result, err := repo.academyRepo.UpdateWeapon(academyWeapon)
	if err != nil {
		return genshin_models.Weapon{}, err
	}

	return result.Weapon, nil
}
