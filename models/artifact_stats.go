package models

type ArtifactStats struct {
	ID                int    `json:ignore`
	OwnerID           int    `json:"OwnerID"`
	Attack            int    `json:"ATK"`
	AttackPercentage  int    `json:"ATK_P"`
	Health            int    `json:"HP"`
	HealthPercentage  int    `json:"HP_P"`
	Defense           int    `json:"DEF"`
	DefensePercentage int    `json:"DEF_P"`
	ElementalMastery  int    `json:"EM"`
	EnergyRecharge    int    `json:"ER"`
	ElementalDamage   int    `json:"ED"`
	CritRate          int    `json:"CR"`
	CritDamage        int    `json:"CD"`
	PhysicalDamage    int    `json:"PHYS"`
	Heal              int    `json:"HEAL"`
	Type              string `json:"TYPE"`
}
