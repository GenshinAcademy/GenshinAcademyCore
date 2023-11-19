package entity

import (
	"ga/internal/types"
	"gorm.io/gorm"
)

type News struct {
	gorm.Model
	Title       types.LocalizedString `gorm:"type:jsonb"`
	Description types.LocalizedString `gorm:"type:jsonb"`
	PreviewUrl  types.LocalizedString `gorm:"type:jsonb"`
	RedirectUrl types.LocalizedString `gorm:"type:jsonb"`
}
