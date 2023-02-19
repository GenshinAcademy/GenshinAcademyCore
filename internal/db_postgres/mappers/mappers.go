// Contains sets of converters from DB academy_models to Core academy_models and from Core academy_models to DB academy_models
package db_mappers

import (
	academy_models "ga/internal/academy_core/models"
	"ga/internal/academy_core/value_objects/localized_string"
	"ga/internal/db_postgres/cache"
	db_models "ga/internal/db_postgres/models"
	genshin_models "ga/pkg/genshin_core/models"
)

type Mapper struct {
	cache        *cache.Cache
	languge      academy_models.Language
	languageName string
}

func CreateMapper(languageName string, language academy_models.Language, cache *cache.Cache) Mapper {
	return Mapper{
		languageName: languageName,
		languge:      language,
		cache:        cache,
	}
}

// MapDbCharacterFromModel Converts DB character model to Core character model
func (mapper Mapper) MapDbCharacterFromModel(model *academy_models.Character) db_models.DbCharacter {
	var strings = mapper.cache.GetCharacterStrings(db_models.DBKey(model.Id))
	var result = db_models.DbCharacter{
		Id:          db_models.DBKey(model.Id),
		CharacterId: db_models.GenshinKey(model.CharacterId),
		Name:        mapper.MapDbStringFromString(strings.Name, model.Name),
		FullName:    mapper.MapDbStringFromString(strings.FullName, model.FullName),
		Description: mapper.MapDbStringFromString(strings.Description, model.Description),
		Title:       mapper.MapDbStringFromString(strings.Title, model.Title),
		Element:     uint8(model.Element),
		Rarity:      uint8(model.Rarity),
		Weapon:      uint8(model.Weapon),
	}

	return result
}

// MapAcademyCharacterFromDbModel Converts Db character model to Academy character model
func (mapper Mapper) MapAcademyCharacterFromDbModel(model *db_models.DbCharacter) academy_models.Character {
	return academy_models.Character{
		AcademyModel: academy_models.AcademyModel{
			Id: academy_models.AcademyId(model.Id),
		},
		Character: mapper.MapGenshinCharacterFromDbModel(model),
	}
}

// MapGenshinCharacterFromDbModel Converts Db character model to Core character model
func (mapper Mapper) MapGenshinCharacterFromDbModel(model *db_models.DbCharacter) genshin_models.Character {
	return genshin_models.Character{
		CharacterId: string(model.CharacterId),
		Name:        mapper.StringFromDbModel(&model.Name),
		FullName:    mapper.StringFromDbModel(&model.FullName),
		Description: mapper.StringFromDbModel(&model.Description),
		Title:       mapper.StringFromDbModel(&model.Title),
		Element:     genshin_models.Element(model.Element),
		Rarity:      genshin_models.Rarity(model.Rarity),
		Weapon:      genshin_models.WeaponType(model.Weapon),
	}
}

// LanguageFromDbModel Converts DB language model to language model
func (mapper Mapper) LanguageFromDbModel(model *db_models.DbLanguage) academy_models.Language {
	return academy_models.Language{
		AcademyModel: academy_models.AcademyModel{
			Id: academy_models.AcademyId(model.Id),
		},
		LanguageName: model.Name,
	}
}

// DbLanguageFromModel Converts language model to DB language model
func (mapper Mapper) DbLanguageFromModel(model *academy_models.Language) db_models.DbLanguage {
	return db_models.DbLanguage{
		Id:   db_models.DBKey(model.Id),
		Name: model.LanguageName,
	}
}

// StringFromDbModel Gets string value from Db model
func (mapper Mapper) StringFromDbModel(model *db_models.DbString) string {
	if model.GetValue() == "" {
		return ""
	}
	return model.GetValue()
}

func (mapper Mapper) MapDbStringFromString(key db_models.DBKey, value string) db_models.DbString {
	return db_models.DbString{
		Id: db_models.DBKey(key),
		StringValues: []db_models.DbStringvalue{
			{
				Id:         db_models.DBKey(key),
				LanguageId: db_models.DBKey(mapper.languge.Id),
				Value:      value,
			},
		},
	}
}

// LocalizedStringFromDbModel Converts DB string model to LocalizedString
func (mapper Mapper) LocalizedStringFromDbModel(model *db_models.DbString) localized_string.LocalizedString {
	if model.GetValue() == "" {
		return localized_string.Empty(localized_string.StringId(model.Id))
	}
	return localized_string.Create(
		localized_string.StringId(model.Id),
		localized_string.LanguageId(model.GetLanguageId()),
		model.GetValue(),
	)
}

// MapDbStringFromLocalizedString Converts LocalizedString to DB string model
func (mapper Mapper) MapDbStringFromLocalizedString(model *localized_string.LocalizedString) db_models.DbString {
	return db_models.DbString{
		Id: db_models.DBKey(model.GetId()),
		StringValues: []db_models.DbStringvalue{
			{
				Id:         db_models.DBKey(model.GetId()),
				LanguageId: db_models.DBKey(model.GetLanguageId()),
				Value:      model.GetValue(),
			},
		},
	}
}
