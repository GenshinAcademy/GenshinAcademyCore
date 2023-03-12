package artifact_proft

import "fmt"

type ProfitSlot string
type ProfitSlotNumber uint8
type StatProfit uint16

const DefaultProfitValue = 0

const (
	SubStatsNumber ProfitSlotNumber = iota
	FlowerNumber
	PlumeNumber
	SandsNumber
	GobletNumber
	CircletNumber
)

const (
	SubStats ProfitSlot = "substats"
	Flower   ProfitSlot = "flower"
	Plume    ProfitSlot = "plume"
	Sands    ProfitSlot = "sands"
	Goblet   ProfitSlot = "goblet"
	Circlet  ProfitSlot = "circlet"
)

type ArtifactProfit struct {
	Slot              ProfitSlot
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

func CreateNew(slot ProfitSlot) ArtifactProfit {
	return ArtifactProfit{
		Slot:              slot,
		Attack:            DefaultProfitValue,
		AttackPercentage:  DefaultProfitValue,
		Health:            DefaultProfitValue,
		HealthPercentage:  DefaultProfitValue,
		Defense:           DefaultProfitValue,
		DefensePercentage: DefaultProfitValue,
		ElementalMastery:  DefaultProfitValue,
		EnergyRecharge:    DefaultProfitValue,
		ElementalDamage:   DefaultProfitValue,
		CritRate:          DefaultProfitValue,
		CritDamage:        DefaultProfitValue,
		PhysicalDamage:    DefaultProfitValue,
		Heal:              DefaultProfitValue,
	}
}

func ProfitSlotFromNumber(num ProfitSlotNumber) ProfitSlot {
	switch num {
	case SubStatsNumber:
		return SubStats
	case FlowerNumber:
		return Flower
	case PlumeNumber:
		return Plume
	case SandsNumber:
		return Sands
	case GobletNumber:
		return Goblet
	case CircletNumber:
		return Circlet
	default:
		// TODO Panic to error return
		panic(fmt.Sprintf("Provided unknown ProfitSlotNumber %d", num))
	}
}

func ProfitSlotNumberFromName(num ProfitSlot) ProfitSlotNumber {
	switch num {
	case SubStats:
		return SubStatsNumber
	case Flower:
		return FlowerNumber
	case Plume:
		return PlumeNumber
	case Sands:
		return SandsNumber
	case Goblet:
		return GobletNumber
	case Circlet:
		return CircletNumber
	default:
		// TODO Panic to error return
		panic(fmt.Sprintf("Provided unknown ProfitSlot %s", num))
	}
}
