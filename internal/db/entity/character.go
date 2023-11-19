package entity

import (
	"ga/internal/types"
	"gorm.io/gorm"
)

type Character struct {
	Id              types.CharacterId     `gorm:"primaryKey"`
	Name            types.LocalizedString `gorm:"type:jsonb"`
	Description     types.LocalizedString `gorm:"type:jsonb"`
	Element         types.Element         `gorm:"type:element"`
	Rarity          types.Rarity          `gorm:"type:rarity"`
	WeaponType      types.WeaponType      `gorm:"type:weapon_type"`
	Icons           CharacterIcons
	ArtifactProfits ArtifactProfits
}

func (c *Character) BeforeCreate(_ *gorm.DB) (err error) {
	if c.ArtifactProfits.CharacterId == "" {
		c.ArtifactProfits.CharacterId = c.Id
	}

	if c.ArtifactProfits.Profits == nil {
		c.ArtifactProfits.Profits = types.DefaultCharacterArtifactProfits()
	}

	return
}
