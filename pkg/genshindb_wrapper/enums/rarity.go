package enums

type Rarity uint

const (
	Legendary Rarity = 5
	Epic      Rarity = 4
	Rare      Rarity = 3
)

type QualityType string

const (
	QualityOrange   QualityType = "QUALITY_ORANGE"
	QualityOrangeSP QualityType = "QUALITY_ORANGE_SP"
	QualityPurple   QualityType = "QUALITY_PURPLE"
)
