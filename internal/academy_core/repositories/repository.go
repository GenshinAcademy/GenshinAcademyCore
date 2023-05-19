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
	GetCharacterIds(parameters find_parameters.CharacterFindParameters) ([]genshin_models.ModelId, error)
	FindCharacterById(characterId models.AcademyId) (models.Character, error)
	FindCharacterByGenshinId(characterId genshin_models.ModelId) (models.Character, error)
	FindCharacters(parameters find_parameters.CharacterFindParameters) ([]models.Character, error)
	AddCharacter(character models.Character) (models.Character, error)
	UpdateCharacter(character models.Character) (models.Character, error)
}

type INewsRepository interface {
	IRepository
	FindNewsById(id models.AcademyId) (models.News, error)
	FindNews(parameters find_parameters.NewsFindParameters) ([]models.News, error)
	AddNews(news models.News) (models.News, error)
	UpdateNews(news models.News) (models.News, error)
}

type ITableRepository interface {
	IRepository
	FindTableById(id models.AcademyId) (models.Table, error)
	FindTables(parameters find_parameters.TableFindParameters) ([]models.Table, error)
	AddTable(table models.Table) (models.Table, error)
	UpdateTable(table models.Table) (models.Table, error)
}

type IWeaponRepository interface {
	IRepository
	GetWeaponIds(parameters find_parameters.WeaponFindParameters) ([]genshin_models.ModelId, error)
	FindWeaponById(id models.AcademyId) (models.Weapon, error)
	FindWeaponByGenshinId(characterId genshin_models.ModelId) (models.Weapon, error)
	FindWeapons(parameters find_parameters.WeaponFindParameters) ([]models.Weapon, error)
	AddWeapon(weapon models.Weapon) (models.Weapon, error)
	UpdateWeapon(weapon models.Weapon) (models.Weapon, error)
}

type IRepositoryProvider interface {
	GetLanguage() *models.Language
	NewCharacterRepo() ICharacterRepository
	CreateNewsRepo() INewsRepository
	CreateTableRepo() ITableRepository
	CreateWeaponRepo() IWeaponRepository
	//NewCharacterIconRepo() ICharacterIconRepository
}
