package models

import (
	"ga/internal/academy_core/value_objects/artifact_profit"
	url "ga/internal/academy_core/value_objects/url"
	"ga/pkg/genshin_core/models"
	"ga/pkg/genshin_core/models/enums"
)

type WeaselAppraiserCharacter struct {
	CharacterId models.ModelId                   `json:"character_id"`
	Name        string                           `json:"name"`
	Element     enums.Element                    `json:"element"`
	IconUrl     url.Url                          `json:"icon_url"`
	StatsProfit []artifact_profit.ArtifactProfit `json:"stats_profit"`
}
