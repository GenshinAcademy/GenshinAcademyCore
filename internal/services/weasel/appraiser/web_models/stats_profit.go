package web_models

type StatsProfit struct {
	Slot              string `json:"slot"`
	Attack            int    `json:"ATK,omitempty"`
	AttackPercentage  int    `json:"ATK_P,omitempty"`
	Health            int    `json:"HP,omitempty"`
	HealthPercentage  int    `json:"HP_P,omitempty"`
	Defense           int    `json:"DEF,omitempty"`
	DefensePercentage int    `json:"DEF_P,omitempty"`
	ElementalMastery  int    `json:"EM,omitempty"`
	EnergyRecharge    int    `json:"ER,omitempty"`
	ElementalDamage   int    `json:"ELEM,omitempty"`
	CritRate          int    `json:"CR,omitempty"`
	CritDamage        int    `json:"CD,omitempty"`
	PhysicalDamage    int    `json:"PHYS,omitempty"`
	Heal              int    `json:"HEAL,omitempty"`
}
