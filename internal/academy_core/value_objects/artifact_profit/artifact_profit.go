package artifact_proft

import (
	"ga/internal/academy_core/value_objects/localized_string"
)

type StatProfit uint16

type ArtifactProfit struct {
	Slot              localized_string.LocalizedString
	Attack            StatProfit
	AttackPercentage  StatProfit
	Health            StatProfit
	HealthPercentage  StatProfit
	Defense           StatProfit
	DefensePercentage StatProfit
	ElementalMastery  StatProfit
	EnergyRecharge    StatProfit
	ElementalDamage   StatProfit
	CritRate          StatProfit
	CritDamage        StatProfit
	PhysicalDamage    StatProfit
	Heal              StatProfit
}

func CreateNew(slot localized_string.LocalizedString) ArtifactProfit {
	return ArtifactProfit{
		Slot: slot,
	}
}
