package artifact_proft

type ProfitSlot string
type ProfitSlotNumber uint8
type StatProfit uint16

const DefaultProfitValue = 0

const (
	PlumeNumber ProfitSlotNumber = iota
	FlowerNumber
	GobletNumber
	CircletNumber
	SandsNumber
)

const (
	Plume   ProfitSlot = "plume"
	Flower  ProfitSlot = "flower"
	Goblet  ProfitSlot = "goblet"
	Circlet ProfitSlot = "circlet"
	Sands   ProfitSlot = "sands"
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
	case PlumeNumber:
		return Plume
	case FlowerNumber:
		return Flower
	case GobletNumber:
		return Goblet
	case CircletNumber:
		return Circlet
	case SandsNumber:
		return Sands
	default:
		panic("Provided unknown ProfitSlotNumber")
	}
}

func ProfitSlotNumberFromName(num ProfitSlot) ProfitSlotNumber {
	switch num {
	case Plume:
		return PlumeNumber
	case Flower:
		return FlowerNumber
	case Goblet:
		return GobletNumber
	case Circlet:
		return CircletNumber
	case Sands:
		return SandsNumber
	default:
		panic("Provided unknown ProfitSlot")
	}
}
