package db

type ArtifactStats struct {
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
