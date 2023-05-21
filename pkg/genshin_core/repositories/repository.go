package repositories

import (
	"ga/pkg/genshin_core/models"
	"ga/pkg/genshin_core/models/languages"
	"ga/pkg/genshin_core/repositories/find_parameters"
	"ga/pkg/genshin_core/value_objects"
)

type Repository interface {
	GetLanguage() *languages.Language
}

type CharacterRepository interface {
	Repository
	GetCharacterIds(parameters find_parameters.CharacterFindParameters) ([]models.ModelId, error)
	FindCharacterById(characterId models.ModelId) (models.Character, error)
	FindCharacters(parameters find_parameters.CharacterFindParameters) ([]models.Character, error)
	AddCharacter(character models.Character) (models.Character, error)
	UpdateCharacter(character models.Character) (models.Character, error)
}

type WeaponRepository interface {
	Repository
	GetWeaponIds(parameters find_parameters.WeaponFindParameters) ([]models.ModelId, error)
	FindWeaponById(weaponId models.ModelId) (models.Weapon, error)
	FindWeapons(find_parameters.WeaponFindParameters) ([]models.Weapon, error)
	AddWeapon(weapon models.Weapon) (models.Weapon, error)
	UpdateWeapon(weapon models.Weapon) (models.Weapon, error)
}

type CharacterIconRepository interface {
	FindIconsByCharacterId(characterId models.ModelId) []value_objects.CharacterIcon
}

type RepositoryProvider interface {
	GetLanguage() *languages.Language
	NewCharacterRepo() CharacterRepository
	NewWeaponRepo() WeaponRepository
	//NewCharacterIconRepo() ICharacterIconRepository
}
