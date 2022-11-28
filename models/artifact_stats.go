package models

type ArtifactStats struct {
	ID                int `json:"id"`
	Attack            int `json:"attack"`
	AttackPercentage  int `json:"attackPercentage"`
	Health            int `json:"health"`
	HealthPercentage  int `json:"healthPercentage"`
	Defense           int `json:"defense"`
	DefensePercentage int `json:"defensePercentage`
	ElementalMastery  int `json:"elementalMastery"`
	EnergyRecharge    int `json:"energyRecharge"`
	ElementalDamage   int `json:"elementalDamage"`
	CritRate          int `json:"critRate"`
	CritDamage        int `json:"critDamage"`
	PhysicalDamage    int `json:"physicalDamage"`
	Heal              int `json:"heal"`
}
