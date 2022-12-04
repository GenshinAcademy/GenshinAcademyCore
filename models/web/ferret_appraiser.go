package web

type CharacterArtifactStatsProfit struct {
	Character
	StatsProfit StatsProfit `json:"statsProfit"`
}

type StatsProfit struct {
	Flower   Flower   `json:"flower"`
	Feather  Feather  `json:"feather"`
	Sands    Sands    `json:"sands"`
	Goblet   Goblet   `json:"goblet"`
	Circlet  Circlet  `json:"circlet"`
	Substats Substats `json:"substats"`
}
