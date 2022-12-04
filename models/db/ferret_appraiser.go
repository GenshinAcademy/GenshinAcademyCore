package db

type StatsProfit struct {
	ID                int `json:"-"`
	OwnerID           int `json:"-"`
	SlotID            int `json:"-"`
	Slot              Slot
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
