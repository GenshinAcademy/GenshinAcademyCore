package models

type CharacterArtifactStatsProfit struct {
	Character
	Flower   Flower   `json:"flower"`
	Feather  Feather  `json:"feather"`
	Sands    Sands    `json:"sands"`
	Goblet   Goblet   `json:"goblet"`
	Circlet  Circlet  `json:"circlet"`
	Substats Substats `json:"substats"`
}
