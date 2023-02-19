package web_models

type FerretCharacter struct {
	CharacterId string        `json:"character_id"`
	Name        string        `json:"name"`
	Element     string        `json:"element"`
	IconUrl     string        `json:"icon_url"`
	StatsProfit []StatsProfit `json:"stats_profit"`
}
