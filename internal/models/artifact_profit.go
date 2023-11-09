package models

import (
	"ga/internal/types"
)

type ArtifactProfit struct {
	Attack            types.StatProfit
	AttackPercentage  types.StatProfit
	Health            types.StatProfit
	HealthPercentage  types.StatProfit
	Defense           types.StatProfit
	DefensePercentage types.StatProfit
	ElementalMastery  types.StatProfit
	EnergyRecharge    types.StatProfit
	ElementalDamage   types.StatProfit
	CritRate          types.StatProfit
	CritDamage        types.StatProfit
	PhysicalDamage    types.StatProfit
	Heal              types.StatProfit
}
