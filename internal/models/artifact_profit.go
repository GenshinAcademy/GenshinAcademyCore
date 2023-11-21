package models

import (
	"ga/internal/types"
)

type ArtifactProfit struct {
	Attack            types.StatProfit `json:"ATK" extensions:"x-order=1"`
	AttackPercentage  types.StatProfit `json:"ATK_P" extensions:"x-order=2"`
	Health            types.StatProfit `json:"HP" extensions:"x-order=3"`
	HealthPercentage  types.StatProfit `json:"HP_P" extensions:"x-order=4"`
	Defense           types.StatProfit `json:"DEF" extensions:"x-order=5"`
	DefensePercentage types.StatProfit `json:"DEF_P" extensions:"x-order=6"`
	ElementalMastery  types.StatProfit `json:"EM" extensions:"x-order=7"`
	EnergyRecharge    types.StatProfit `json:"ER" extensions:"x-order=8"`
	ElementalDamage   types.StatProfit `json:"ELEM" extensions:"x-order=9"`
	CritRate          types.StatProfit `json:"CR" extensions:"x-order=10"`
	CritDamage        types.StatProfit `json:"CD" extensions:"x-order=11"`
	PhysicalDamage    types.StatProfit `json:"PHYS" extensions:"x-order=12"`
	Heal              types.StatProfit `json:"HEAL" extensions:"x-order=13"`
}
