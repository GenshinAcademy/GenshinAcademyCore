package models

import (
	"ga/internal/academy_core/value_objects/artifact_profit"
	url "ga/internal/academy_core/value_objects/url"
	"ga/pkg/genshin_core/models"
	"ga/pkg/genshin_core/models/enums"
)

type WeaselAppraiserCharacter struct {
	CharacterId models.ModelId                   `json:"character_id" example:"lisa" extensions:"x-order=0"`
	Name        string                           `json:"name" example:"Lisa" extensions:"x-order=1"`
	Element     enums.Element                    `json:"element" example:"5" extensions:"x-order=2"`
	IconUrl     url.Url                          `json:"icon_url" example:"https://example.com" extensions:"x-order=3"`
	StatsProfit []artifact_profit.ArtifactProfit `json:"stats_profit" extensions:"x-order=4"`
} //@name WeaselAppraiserCharacter
