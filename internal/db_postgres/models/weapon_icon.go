package db_models

// WeaponIcon represents table for Genshin weapons' icons in database.
type WeaponIcon struct {
	WeaponId DBKey  `gorm:"primaryKey"`
	IconType uint8  `gorm:"primaryKey"`
	Url      string `gorm:"type:varchar"`
}

// TODO: Make URL type that checks format
