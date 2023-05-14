package repositories

import (
	"ga/internal/academy_core/models"
	"ga/internal/academy_core/repositories/find_parameters"
	genshin_models "ga/pkg/genshin_core/models"
	"ga/pkg/genshin_core/models/languages"
)

type IRepository interface {
	GetLanguage() *models.Language
}

type ILanguageRepository interface {
	// AddLanguage adds language to the database
	AddLanguage(language *models.Language)

	// FindLanguage finds language in database.
	// If none is found returned ID is 0.
	FindLanguage(lang *languages.Language) *models.Language
}

type ICharacterRepository interface {
	IRepository
	GetCharacterIds(parameters find_parameters.CharacterFindParameters) []genshin_models.ModelId
	FindCharacterById(characterId models.AcademyId) (*models.Character, bool)
	FindCharacterByGenshinId(characterId genshin_models.ModelId) (*models.Character, bool)
	FindCharacters(parameters find_parameters.CharacterFindParameters) []models.Character
	AddCharacter(character *models.Character) (*models.Character, error)
	UpdateCharacter(character *models.Character) (*models.Character, error)
}

type INewsRepository interface {
	IRepository
	FindNewsById(id models.AcademyId) *models.News
	FindNews(parameters find_parameters.NewsFindParameters) []models.News
	AddNews(news *models.News) (*models.News, error)
	UpdateNews(news *models.News) (*models.News, error)
}

type ITableRepository interface {
	IRepository
	FindTableById(id models.AcademyId) *models.Table
	FindTables(parameters find_parameters.TableFindParameters) []models.Table
	AddTable(table *models.Table) (*models.Table, error)
	UpdateTable(table *models.Table) (*models.Table, error)
}

type IWeaponRepository interface {
	IRepository
	GetWeaponIds(parameters find_parameters.WeaponFindParameters) []genshin_models.ModelId
	FindWeaponById(id models.AcademyId) *models.Weapon
	FindWeaponByGenshinId(characterId genshin_models.ModelId) (*models.Weapon, bool)
	FindWeapons(parameters find_parameters.WeaponFindParameters) []models.Weapon
	AddWeapon(weapon *models.Weapon) (*models.Weapon, error)
	UpdateWeapon(weapon *models.Weapon) (*models.Weapon, error)
}

type IRepositoryProvider interface {
	GetLanguage() *models.Language
	NewCharacterRepo() ICharacterRepository
	CreateNewsRepo() INewsRepository
	CreateTableRepo() ITableRepository
	CreateWeaponRepo() IWeaponRepository
	//NewCharacterIconRepo() ICharacterIconRepository
}
