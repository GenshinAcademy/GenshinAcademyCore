package web_models

import url "ga/internal/academy_core/value_objects/url"

type WeaselAppraiserCharacter struct {
	CharacterId string        `json:"character_id"`
	Name        string        `json:"name"`
	Element     uint8         `json:"element"`
	IconUrl     url.Url       `json:"icon_url"`
	StatsProfit []StatsProfit `json:"stats_profit"`
}
