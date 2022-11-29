package models

type CharacterArtifactStatsProfit struct {
	ID int `json:"id" gorm:"primary_key"`
	Character
	Flower   Flower   `json:"flower" gorm:"foreignKey:OwnerID"`
	Feather  Feather  `json:"feather" gorm:"foreignKey:OwnerID"`
	Sands    Sands    `json:"sands" gorm:"foreignKey:OwnerID"`
	Goblet   Goblet   `json:"goblet" gorm:"foreignKey:OwnerID"`
	Circlet  Circlet  `json:"circlet" gorm:"foreignKey:OwnerID"`
	Substats Substats `json:"substats" gorm:"foreignKey:OwnerID"`
}
