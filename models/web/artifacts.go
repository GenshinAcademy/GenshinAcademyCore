package web

type Flower struct {
	Health int `json:"HP"`
}

type Feather struct {
	Attack int `json:"ATK"`
}

type Sands struct {
	AttackPercentage  int `json:"ATK_P"`
	HealthPercentage  int `json:"HP_P"`
	DefensePercentage int `json:"DEF_P"`
	ElementalMastery  int `json:"EM"`
	EnergyRecharge    int `json:"ER"`
}

type Goblet struct {
	AttackPercentage  int `json:"ATK_P"`
	HealthPercentage  int `json:"HP_P"`
	DefensePercentage int `json:"DEF_P"`
	ElementalMastery  int `json:"EM"`
	PhysicalDamage    int `json:"PHYS"`
	ElementalDamage   int `json:"ELEM"`
}

type Circlet struct {
	AttackPercentage  int `json:"ATK_P"`
	DefensePercentage int `json:"DEF_P"`
	HealthPercentage  int `json:"HP_P"`
	ElementalMastery  int `json:"EM"`
	CritRate          int `json:"CR"`
	CritDamage        int `json:"CD"`
	Heal              int `json:"HEAL"`
}

type Substats struct {
	Attack            int `json:"ATK"`
	AttackPercentage  int `json:"ATK_P"`
	Health            int `json:"HP"`
	HealthPercentage  int `json:"HP_P"`
	CritDamage        int `json:"CD"`
	CritRate          int `json:"CR"`
	ElementalMastery  int `json:"EM"`
	Defense           int `json:"DEF"`
	DefensePercentage int `json:"DEF_P"`
	EnergyRecharge    int `json:"ER"`
}
