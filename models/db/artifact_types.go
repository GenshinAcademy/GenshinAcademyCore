package db

import (
	models "genshinacademycore/models/web"
)

type FlowerProfit struct {
	ID int64
	models.Flower
}

type FeatherProfit struct {
	ID int64
	models.Feather
}

type SandsProfit struct {
	ID int64
	models.Sands
}

type GobletProfit struct {
	ID int64
	models.Goblet
}

type CircletProfit struct {
	ID int64
	models.Circlet
}

type SubstatsProfit struct {
	ID int64
	models.Substats
}
