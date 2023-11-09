package entity

import (
	"ga/internal/types"
)

type ArtifactProfits struct {
	CharacterId types.CharacterId              `gorm:"primaryKey"`
	Profits     types.CharacterArtifactProfits `gorm:"type:jsonb"`
}
