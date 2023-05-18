// Package db_mappers Contains sets of converters from DB academy_models to Core academy_models and from Core academy_models to DB academy_models
package db_mappers

import (
	// Academy specific imports

	academy_models "ga/internal/academy_core/models"
	artifact_profit "ga/internal/academy_core/value_objects/artifact_profit"
	localized_string "ga/internal/academy_core/value_objects/localized_string"

	// Genshin specific imports
	genshin_models "ga/pkg/genshin_core/models"
	genshin_enums "ga/pkg/genshin_core/models/enums"
	genshin_objects "ga/pkg/genshin_core/value_objects"

	// Misc
	"ga/internal/db_postgres/cache"
	db_models "ga/internal/db_postgres/models"
)

type Mapper struct {
	cache        *cache.Cache
	language     *academy_models.Language
	languageName string
}

func CreateMapper(languageName string, language *academy_models.Language, cache *cache.Cache) Mapper {
	return Mapper{
		languageName: languageName,
		language:     language,
		cache:        cache,
	}
}

// MapDbCharacterFromModel converts Core character model to DB character model.
func (mapper Mapper) MapDbCharacterFromModel(model academy_models.Character) db_models.Character {
	var strings = mapper.cache.GetCharacterStrings(db_models.DBKey(model.Id))

	var character = db_models.Character{
		Id:          db_models.DBKey(model.Id),
		CharacterId: db_models.GenshinKey(model.Character.Id),
		Name:        *mapper.MapDbStringFromString(strings.Name, model.Name),
		FullName:    *mapper.MapDbStringFromString(strings.FullName, model.FullName),
		Description: *mapper.MapDbStringFromString(strings.Description, model.Description),
		Title:       *mapper.MapDbStringFromString(strings.Title, model.Title),
		Element:     uint8(model.Element),
		Rarity:      uint8(model.Rarity),
		Weapon:      uint8(model.Weapon),
	}
	mapper.mapDbCharacterArrays(&character, model)

	return character
}

// MapAcademyCharacterFromDbModel converts DB character model to Academy character model.
func (mapper Mapper) MapAcademyCharacterFromDbModel(model db_models.Character) academy_models.Character {
	var character = academy_models.Character{
		AcademyModel: academy_models.AcademyModel{
			Id: academy_models.AcademyId(model.Id),
		},
		Character: mapper.MapGenshinCharacterFromDbModel(model),
	}
	mapper.mapAcademyCharacterArrays(&character, model)
	mapper.cache.UpdateCharacterStrings(model)

	return character
}

func (mapper Mapper) mapDbCharacterArrays(model *db_models.Character, srcModel academy_models.Character) {
	// Genshin related
	for i := 0; i < len(srcModel.Icons); i += 1 {
		model.Icons = append(model.Icons, mapper.MapDbCharacterIcon(db_models.DBKey(srcModel.Id), srcModel.Icons[i]))
	}

	//Academy related
	for i := 0; i < len(srcModel.Profits); i += 1 {
		model.ArtifactProfits = append(model.ArtifactProfits, *mapper.MapDbArtifactStats(db_models.DBKey(srcModel.Id), &srcModel.Profits[i]))
	}
}

func (mapper Mapper) mapAcademyCharacterArrays(model *academy_models.Character, srcModel db_models.Character) {
	for i := 0; i < len(srcModel.ArtifactProfits); i += 1 {
		model.Profits = append(model.Profits, mapper.MapArtifactStats(srcModel.ArtifactProfits[i]))
	}
}

func (mapper Mapper) mapGenshinCharacterArrays(model *genshin_models.Character, srcModel db_models.Character) {
	for i := 0; i < len(srcModel.Icons); i += 1 {
		model.Icons = append(model.Icons, mapper.MapCharacterIcon(srcModel.Icons[i]))
	}
}

// MapGenshinCharacterFromDbModel converts DB character model to Core character model
func (mapper Mapper) MapGenshinCharacterFromDbModel(model db_models.Character) genshin_models.Character {
	var character = genshin_models.Character{
		BaseModel: genshin_models.BaseModel{
			Id: genshin_models.ModelId(model.CharacterId),
		},
		Name:        mapper.StringFromDbModel(model.Name),
		FullName:    mapper.StringFromDbModel(model.FullName),
		Description: mapper.StringFromDbModel(model.Description),
		Title:       mapper.StringFromDbModel(model.Title),
		Element:     genshin_enums.Element(model.Element),
		Rarity:      genshin_enums.Rarity(model.Rarity),
		Weapon:      genshin_enums.WeaponType(model.Weapon),
	}
	mapper.mapGenshinCharacterArrays(&character, model)
	mapper.cache.UpdateCharacterStrings(model)

	return character
}

// MapLanguageFromDbModel converts DB language model to language model
func (mapper Mapper) MapLanguageFromDbModel(model *db_models.Language) *academy_models.Language {
	return &academy_models.Language{
		AcademyModel: academy_models.AcademyModel{
			Id: academy_models.AcademyId(model.Id),
		},
		LanguageName: model.Name,
	}
}

// DbLanguageFromModel converts language model to DB language model
func (mapper Mapper) DbLanguageFromModel(model *academy_models.Language) *db_models.Language {
	return &db_models.Language{
		Id:   db_models.DBKey(model.Id),
		Name: model.LanguageName,
	}
}

// StringFromDbModel returns string value from DB model
func (mapper Mapper) StringFromDbModel(model db_models.String) string {
	if model.GetValue() == "" {
		return ""
	}
	return model.GetValue()
}

func (mapper Mapper) MapDbStringFromString(key db_models.DBKey, value string) *db_models.String {
	return &db_models.String{
		Id: db_models.DBKey(key),
		StringValues: []db_models.StringValue{
			{
				Id:         db_models.DBKey(key),
				LanguageId: db_models.DBKey(mapper.language.Id),
				Value:      value,
			},
		},
	}
}

// LocalizedStringFromDbModel Converts DB string model to LocalizedString
func (mapper Mapper) LocalizedStringFromDbModel(model *db_models.String) *localized_string.LocalizedString {
	var result localized_string.LocalizedString
	if model.GetValue() == "" {
		result = localized_string.Empty(localized_string.StringId(model.Id))
		return &result
	}

	result = localized_string.Create(
		localized_string.StringId(model.Id),
		localized_string.LanguageId(model.GetLanguageId()),
		model.GetValue(),
	)

	return &result
}

// MapDbStringFromLocalizedString Converts LocalizedString to DB string model
func (mapper Mapper) MapDbStringFromLocalizedString(model *localized_string.LocalizedString) *db_models.String {
	return &db_models.String{
		Id: db_models.DBKey(model.GetId()),
		StringValues: []db_models.StringValue{
			{
				Id:         db_models.DBKey(model.GetId()),
				LanguageId: db_models.DBKey(model.GetLanguageId()),
				Value:      model.GetValue(),
			},
		},
	}
}

func (mapper Mapper) MapArtifactStats(model db_models.ArtifactProfit) artifact_profit.ArtifactProfit {
	return artifact_profit.ArtifactProfit{
		Slot:              artifact_profit.ProfitSlotFromNumber(artifact_profit.ProfitSlotNumber(model.SlotId)),
		Attack:            artifact_profit.StatProfit(model.Attack),
		AttackPercentage:  artifact_profit.StatProfit(model.AttackPercentage),
		Health:            artifact_profit.StatProfit(model.Health),
		HealthPercentage:  artifact_profit.StatProfit(model.HealthPercentage),
		Defense:           artifact_profit.StatProfit(model.Defense),
		DefensePercentage: artifact_profit.StatProfit(model.DefensePercentage),
		ElementalMastery:  artifact_profit.StatProfit(model.ElementalMastery),
		EnergyRecharge:    artifact_profit.StatProfit(model.EnergyRecharge),
		ElementalDamage:   artifact_profit.StatProfit(model.ElementalDamage),
		CritRate:          artifact_profit.StatProfit(model.CritRate),
		CritDamage:        artifact_profit.StatProfit(model.CritDamage),
		PhysicalDamage:    artifact_profit.StatProfit(model.PhysicalDamage),
		Heal:              artifact_profit.StatProfit(model.Heal),
	}
}

func (mapper Mapper) MapDbArtifactStats(parentId db_models.DBKey, model *artifact_profit.ArtifactProfit) *db_models.ArtifactProfit {
	if parentId == db_models.DBKey(academy_models.UNDEFINED_ID) {
		panic("Cannot create artifactstat with undefined character")
	}
	return &db_models.ArtifactProfit{
		CharacterId:       parentId,
		SlotId:            db_models.DBKey(artifact_profit.ProfitSlotNumberFromName(model.Slot)),
		Attack:            uint16(model.Attack),
		AttackPercentage:  uint16(model.AttackPercentage),
		Health:            uint16(model.Health),
		HealthPercentage:  uint16(model.HealthPercentage),
		Defense:           uint16(model.Defense),
		DefensePercentage: uint16(model.DefensePercentage),
		ElementalMastery:  uint16(model.ElementalMastery),
		EnergyRecharge:    uint16(model.EnergyRecharge),
		ElementalDamage:   uint16(model.ElementalDamage),
		CritRate:          uint16(model.CritRate),
		CritDamage:        uint16(model.CritDamage),
		PhysicalDamage:    uint16(model.PhysicalDamage),
		Heal:              uint16(model.Heal),
	}
}

func (mapper Mapper) MapCharacterIcon(model db_models.CharacterIcon) genshin_objects.CharacterIcon {
	return genshin_objects.CharacterIcon{
		Type: genshin_objects.CharacterIconType(model.IconType),
		Url:  model.Url,
	}
}

func (mapper Mapper) MapDbCharacterIcon(parentId db_models.DBKey, model genshin_objects.CharacterIcon) db_models.CharacterIcon {
	return db_models.CharacterIcon{
		CharacterId: parentId,
		IconType:    uint8(model.Type),
		Url:         model.Url,
	}
}

func (mapper Mapper) MapNewsFromDbModel(model *db_models.News) *academy_models.News {
	var modelNews = academy_models.News{
		AcademyModel: academy_models.AcademyModel{
			Id: academy_models.AcademyId(model.Id),
		},
		Title:       mapper.StringFromDbModel(model.Title),
		Description: mapper.StringFromDbModel(model.Description),
		Preview:     mapper.StringFromDbModel(model.PreviewUrl),
		RedirectUrl: mapper.StringFromDbModel(model.RedirectUrl),
		CreatedAt:   model.CreatedAt,
	}
	mapper.cache.UpdateNewsStrings(model)

	return &modelNews
}

func (mapper Mapper) MapDbNewsFromModel(model *academy_models.News) *db_models.News {
	var newsStrings = mapper.cache.GetNewsStrings(db_models.DBKey(model.Id))

	return &db_models.News{
		Id:          db_models.DBKey(model.Id),
		Title:       *mapper.MapDbStringFromString(newsStrings.Title, model.Title),
		Description: *mapper.MapDbStringFromString(newsStrings.Description, model.Description),
		PreviewUrl:  *mapper.MapDbStringFromString(newsStrings.Preview, model.Preview),
		RedirectUrl: *mapper.MapDbStringFromString(newsStrings.Redirect, model.RedirectUrl),
		CreatedAt:   model.CreatedAt,
	}
}

func (mapper Mapper) MapTableFromDbModel(model *db_models.Table) *academy_models.Table {
	var modelNews = academy_models.Table{
		AcademyModel: academy_models.AcademyModel{
			Id: academy_models.AcademyId(model.Id),
		},
		Title:       mapper.StringFromDbModel(model.Title),
		Description: mapper.StringFromDbModel(model.Description),
		Icon:        model.PreviewUrl,
		RedirectUrl: mapper.StringFromDbModel(model.RedirectUrl),
	}
	mapper.cache.UpdateTableStrings(model)

	return &modelNews
}

func (mapper Mapper) MapDbTableFromModel(model *academy_models.Table) *db_models.Table {
	var tableStrings = mapper.cache.GetTableStrings(db_models.DBKey(model.Id))

	return &db_models.Table{
		Id:          db_models.DBKey(model.Id),
		Title:       *mapper.MapDbStringFromString(tableStrings.Title, model.Title),
		Description: *mapper.MapDbStringFromString(tableStrings.Description, model.Description),
		PreviewUrl:  string(model.Icon),
		RedirectUrl: *mapper.MapDbStringFromString(tableStrings.Redirect, model.RedirectUrl),
	}
}

func (mapper Mapper) MapWeaponIcon(model *db_models.WeaponIcon) *genshin_objects.WeaponIcon {
	return &genshin_objects.WeaponIcon{
		Type: genshin_objects.WeaponIconType(model.IconType),
		Url:  model.Url,
	}
}

func (mapper Mapper) MapDbWeaponIcon(parentId db_models.DBKey, model *genshin_objects.WeaponIcon) *db_models.WeaponIcon {
	return &db_models.WeaponIcon{
		WeaponId: parentId,
		IconType: uint8(model.Type),
		Url:      model.Url,
	}
}

func (mapper Mapper) MapDbWeaponFromModel(model *academy_models.Weapon) *db_models.Weapon {
	var strings = mapper.cache.GetWeaponStrings(db_models.DBKey(model.Id))

	var weapon = &db_models.Weapon{
		Id:                db_models.DBKey(model.Id),
		WeaponId:          db_models.GenshinKey(model.Weapon.Id),
		Name:              *mapper.MapDbStringFromString(strings.Name, model.Name),
		Description:       *mapper.MapDbStringFromString(strings.Description, model.Description),
		DescriptionRaw:    *mapper.MapDbStringFromString(strings.DescriptionRaw, model.DescriptionRaw),
		EffectName:        *mapper.MapDbStringFromString(strings.EffectName, model.EffectName),
		EffectTemplateRaw: *mapper.MapDbStringFromString(strings.EffectTemplateRaw, model.EffectTemplateRaw),
		BaseAttack:        model.BaseAttackValue,
		BaseStatText:      model.BaseStatText,
		MainStatType:      model.MainStatType,
		MainStatName:      model.MainStatName,
		Rarity:            uint8(model.Rarity),
		Type:              uint8(model.WeaponType),
		Icons:             make([]db_models.WeaponIcon, 0),
	}
	mapper.mapDbWeaponArrays(weapon, model)

	return weapon
}

// MapAcademyWeaponFromDbModel converts DB character model to Academy character model.
func (mapper Mapper) MapAcademyWeaponFromDbModel(model *db_models.Weapon) *academy_models.Weapon {
	var weapon = &academy_models.Weapon{
		AcademyModel: academy_models.AcademyModel{
			Id: academy_models.AcademyId(model.Id),
		},
		Weapon: *mapper.MapGenshinWeaponFromDbModel(model),
	}
	mapper.mapAcademyWeaponArrays(weapon, model)
	mapper.cache.UpdateWeaponStrings(model)

	return weapon
}

func (mapper Mapper) mapDbWeaponArrays(model *db_models.Weapon, srcModel *academy_models.Weapon) {
	// Genshin related
	for i := 0; i < len(srcModel.Icons); i += 1 {
		model.Icons = append(model.Icons, *mapper.MapDbWeaponIcon(db_models.DBKey(srcModel.Id), &srcModel.Icons[i]))
	}

	//Academy related
}

func (mapper Mapper) mapAcademyWeaponArrays(model *academy_models.Weapon, srcModel *db_models.Weapon) {
}

func (mapper Mapper) mapGenshinWeaponArrays(model *genshin_models.Weapon, srcModel *db_models.Weapon) {
	for i := 0; i < len(srcModel.Icons); i += 1 {
		model.Icons = append(model.Icons, *mapper.MapWeaponIcon(&srcModel.Icons[i]))
	}
}

// MapGenshinCharacterFromDbModel converts DB character model to Core character model
func (mapper Mapper) MapGenshinWeaponFromDbModel(model *db_models.Weapon) *genshin_models.Weapon {
	var weapon = genshin_models.Weapon{
		BaseModel: genshin_models.BaseModel{
			Id: genshin_models.ModelId(model.WeaponId),
		},
		Name:              mapper.StringFromDbModel(model.Name),
		Description:       mapper.StringFromDbModel(model.Description),
		DescriptionRaw:    mapper.StringFromDbModel(model.DescriptionRaw),
		EffectName:        mapper.StringFromDbModel(model.EffectName),
		EffectTemplateRaw: mapper.StringFromDbModel(model.EffectTemplateRaw),
		BaseAttackValue:   model.BaseAttack,
		BaseStatText:      model.BaseStatText,
		MainStatType:      model.MainStatType,
		MainStatName:      model.MainStatName,
		Rarity:            genshin_enums.Rarity(model.Rarity),
		WeaponType:        genshin_enums.WeaponType(model.Type),
		Icons:             make([]genshin_objects.WeaponIcon, 0),
	}
	mapper.mapGenshinWeaponArrays(&weapon, model)
	mapper.cache.UpdateWeaponStrings(model)

	return &weapon
}
