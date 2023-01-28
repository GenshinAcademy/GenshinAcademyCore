// Contains sets of converters from DB models to Core models and from Core models to DB models
package db_mappers

import (
	db_models "ga/internal/db_postgres/models"
	models "ga/pkg/core/models"
	"ga/pkg/core/value_objects/localized_string"
)

// Converts DB character model to Core character model
func DbCharacterFromModel(model *models.Character) db_models.Db_Character {

	var result db_models.Db_Character = db_models.Db_Character{
		Id:          db_models.DBKey(model.Id),
		CharacterId: model.CharacterId,
		Name:        DbStringFromLocalizedString(&model.Name),
		FullName:    DbStringFromLocalizedString(&model.FullName),
		Description: DbStringFromLocalizedString(&model.Description),
		Title:       DbStringFromLocalizedString(&model.Title),
		Element:     byte(model.Element),
		Rarity:      byte(model.Rarity),
		Weapon:      byte(model.Weapon),
	}

	return result
}

// Converts Core character model to DB character model
func CharacterfromDbModel(model *db_models.Db_Character) models.Character {
	return models.Character{
		BaseModel: models.BaseModel{
			Id: models.ModelId(model.Id),
		},
		CharacterId: model.CharacterId,
		Name:        LocalizedStringFromDbModel(&model.Name),
		FullName:    LocalizedStringFromDbModel(&model.FullName),
		Description: LocalizedStringFromDbModel(&model.Description),
		Title:       LocalizedStringFromDbModel(&model.Title),
		Element:     models.Element(model.Element),
		Rarity:      models.Rarity(model.Rarity),
		Weapon:      models.WeaponType(model.Weapon),
	}
}

// Converts DB language model to Core language model
func LanguageFromDbModel(model *db_models.Db_Language) models.Language {
	return models.Language{
		BaseModel: models.BaseModel{
			Id: models.ModelId(model.Id),
		},
		LanguageName: model.Name,
	}
}

// Converts Core language model to DB language model
func DbLanguageFromModel(model *models.Language) db_models.Db_Language {
	return db_models.Db_Language{
		Id:   db_models.DBKey(model.Id),
		Name: model.LanguageName,
	}
}

// Converts DB string model to LocalizedString
func LocalizedStringFromDbModel(model *db_models.Db_String) localized_string.LocalizedString {
	if model.GetValue() == "" {
		return localized_string.Empty(localized_string.StringId(model.Id))
	}
	return localized_string.Create(
		localized_string.StringId(model.Id),
		localized_string.LanguageId(model.GetLanguageId()),
		model.GetValue(),
	)
}

// Converts LocalizedString to DB string model
func DbStringFromLocalizedString(model *localized_string.LocalizedString) db_models.Db_String {
	return db_models.Db_String{
		Id: db_models.DBKey(model.GetId()),
		StringValues: []db_models.Db_StringValue{
			{
				Id:         db_models.DBKey(model.GetId()),
				LanguageId: db_models.DBKey(model.GetLanguageId()),
				Value:      model.GetValue(),
			},
		},
	}
}
