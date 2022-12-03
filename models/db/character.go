package db

type Character struct {
	ID          int `json:"-"`
	CharacterId string
	Element     string
	Name        Name        `gorm:"foreignKey:ID"`
	StatsProfit StatsProfit `gorm:"foreignKey:ID"`
}

type Name struct {
	ID      int `json:"-"`
	English string
	Russian string
}

type StatsProfit struct {
	ID                int `json:"-"`
	Type              string
	Attack            int
	AttackPercentage  int
	Health            int
	HealthPercentage  int
	Defense           int
	DefensePercentage int
	ElementalMastery  int
	EnergyRecharge    int
	ElementalDamage   int
	CritRate          int
	CritDamage        int
	PhysicalDamage    int
	Heal              int
}

// type StatsProfit struct {
// 	ID       int
// 	Flower   FlowerProfit   `gorm:"foreignKey:ID"`
// 	Feather  FeatherProfit  `gorm:"foreignKey:ID"`
// 	Sands    SandsProfit    `gorm:"foreignKey:ID"`
// 	Goblet   GobletProfit   `gorm:"foreignKey:ID"`
// 	Circlet  CircletProfit  `gorm:"foreignKey:ID"`
// 	Substats SubstatsProfit `gorm:"foreignKey:ID"`
// }
