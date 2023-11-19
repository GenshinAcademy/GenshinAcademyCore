package entity

import (
	"ga/internal/types"
)

type CharacterIcons struct {
	CharacterId types.CharacterId `gorm:"primaryKey"`
	Url         types.IconsUrl    `gorm:"type:jsonb"`
}
