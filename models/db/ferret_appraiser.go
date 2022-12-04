package db

type StatsProfit struct {
	ID                int  `json:"-"`
	OwnerID           int  `json:"-"`
	Slot              Slot `gorm:"many2many:artifact_slots"`
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
