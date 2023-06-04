package enums

type Rarity uint        //@name Rarity
type QualityType string //@name QualityType

const (
	Legendary Rarity = 5
	Epic      Rarity = 4
	Rare      Rarity = 3
)

const (
	QualityOrange   QualityType = "QUALITY_ORANGE"
	QualityOrangeSP QualityType = "QUALITY_ORANGE_SP"
	QualityPurple   QualityType = "QUALITY_PURPLE"
)
