package entity

import (
	"ga/internal/types"
	"gorm.io/gorm"
)

type Table struct {
	gorm.Model
	Title       types.LocalizedString `gorm:"type:jsonb"`
	Description types.LocalizedString `gorm:"type:jsonb"`
	IconUrl     string
	RedirectUrl types.LocalizedString `gorm:"type:jsonb"`
}
