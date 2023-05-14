package db_models

type Weapon struct {
	//Primary properties
	Id       DBKey      `gorm:"primaryKey"`
	WeaponId GenshinKey `gorm:"weapon_id;uniqueIndex;type:varchar"`
	//Localized strings
	NameId              DBKey `gorm:"column:name"`
	Name                String
	DescriptionId       DBKey `gorm:"column:description"`
	Description         String
	DescriptionRawId    DBKey `gorm:"column:description_raw"`
	DescriptionRaw      String
	EffectNameId        DBKey `gorm:"colum:effect_name"`
	EffectName          String
	EffectTemplateRawId DBKey `gorm:"colum:effect_template_raw"`
	EffectTemplateRaw   String
	//Other prooperties
	BaseAttack   float64 `gorm:"base_attack"`
	MainStatType string  `gorm:"colum:main_stat_type;type:varchar"`
	MainStatName string  `gorm:"column:main_stat_name;type:varchar"`
	BaseStatText string  `gorm:"colum:base_stat_text;type:varchar"`
	Rarity       uint8   `gorm:"column:rarity"`
	Type         uint8   `gorm:"column:type"`
	//Related objects
	Icons []WeaponIcon `gorm:"foreignKey:WeaponId"`
}
