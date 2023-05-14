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

func (repo PostgresGenshinWeaponRepository) GetWeaponIds(parameters find_parameters.WeaponFindParameters) []genshin_models.ModelId {
	var academyParams = academy_find_parameters.WeaponFindParameters{
		WeaponFindParameters: parameters,
	}

	return repo.academyRepo.GetWeaponIds(academyParams)
}

func (repo PostgresGenshinWeaponRepository) FindWeaponById(weaponId genshin_models.ModelId) (genshin_models.Weapon, error) {
	var weapon, found = repo.academyRepo.FindWeaponByGenshinId(weaponId)
	if !found {
		return genshin_models.Weapon{}, errors.WeaponNotFound(weaponId)
	}

	return weapon.Weapon, nil
}

func (repo PostgresGenshinWeaponRepository) FindWeapons(parameters find_parameters.WeaponFindParameters) []genshin_models.Weapon {
	var weapons = make([]genshin_models.Weapon, 0)
	var academyParams = academy_find_parameters.WeaponFindParameters{
		WeaponFindParameters: parameters,
	}

	var academyWeapons = repo.academyRepo.FindWeapons(academyParams)
	for i := 0; i < len(academyWeapons); i += 1 {
		weapons = append(weapons, academyWeapons[i].Weapon)
	}

	return weapons
}

func (repo PostgresGenshinWeaponRepository) AddWeapon(model *genshin_models.Weapon) error {
	if model.Id == genshin_models.DEFAULT_ID {
		return errors.EmptyId()
	}

	var academyWeapon, _ = repo.academyRepo.FindWeaponByGenshinId(model.Id)
	academyWeapon.Weapon = *model
	repo.academyRepo.AddWeapon(academyWeapon)

	return nil
}

func (repo PostgresGenshinWeaponRepository) UpdateWeapon(model *genshin_models.Weapon) error {
	if model.Id == genshin_models.DEFAULT_ID {
		return errors.EmptyId()
	}

	var academyWeapon, found = repo.academyRepo.FindWeaponByGenshinId(model.Id)
	if !found {
		return errors.WeaponNotFound(model.Id)
	}

	academyWeapon.Weapon = *model
	repo.academyRepo.UpdateWeapon(academyWeapon)

	return nil
}
