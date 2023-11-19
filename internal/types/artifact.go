package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type ArtifactSlot string

const (
	SubStats ArtifactSlot = "substats"
	Flower   ArtifactSlot = "flower"
	Plume    ArtifactSlot = "plume"
	Sands    ArtifactSlot = "sands"
	Goblet   ArtifactSlot = "goblet"
	Circlet  ArtifactSlot = "circlet"
)

type StatType string

const (
	Attack            StatType = "ATK"
	AttackPercentage  StatType = "ATK_P"
	Health            StatType = "HP"
	HealthPercentage  StatType = "HP_P"
	Defence           StatType = "DEF"
	DefencePercentage StatType = "DEF_P"
	ElementalMastery  StatType = "EM"
	EnergyRecharge    StatType = "ER"
	ElementalDamage   StatType = "ELEM"
	CritRate          StatType = "CR"
	CritDamage        StatType = "CD"
	PhysicalDamage    StatType = "PHYS"
	Heal              StatType = "HEAL"
)

type StatProfit uint16

type ArtifactProfitId uint

type CharacterArtifactProfits map[ArtifactSlot]map[StatType]StatProfit

func DefaultCharacterArtifactProfits() CharacterArtifactProfits {
	return CharacterArtifactProfits{
		SubStats: map[StatType]StatProfit{
			Attack:            0,
			AttackPercentage:  0,
			Health:            0,
			HealthPercentage:  0,
			Defence:           0,
			DefencePercentage: 0,
			ElementalMastery:  0,
			EnergyRecharge:    0,
			ElementalDamage:   0,
			CritRate:          0,
			CritDamage:        0,
			PhysicalDamage:    0,
			Heal:              0,
		},
		Flower: map[StatType]StatProfit{
			Attack:            0,
			AttackPercentage:  0,
			Health:            0,
			HealthPercentage:  0,
			Defence:           0,
			DefencePercentage: 0,
			ElementalMastery:  0,
			EnergyRecharge:    0,
			ElementalDamage:   0,
			CritRate:          0,
			CritDamage:        0,
			PhysicalDamage:    0,
			Heal:              0,
		},
		Plume: map[StatType]StatProfit{
			Attack:            0,
			AttackPercentage:  0,
			Health:            0,
			HealthPercentage:  0,
			Defence:           0,
			DefencePercentage: 0,
			ElementalMastery:  0,
			EnergyRecharge:    0,
			ElementalDamage:   0,
			CritRate:          0,
			CritDamage:        0,
			PhysicalDamage:    0,
			Heal:              0,
		},
		Sands: map[StatType]StatProfit{
			Attack:            0,
			AttackPercentage:  0,
			Health:            0,
			HealthPercentage:  0,
			Defence:           0,
			DefencePercentage: 0,
			ElementalMastery:  0,
			EnergyRecharge:    0,
			ElementalDamage:   0,
			CritRate:          0,
			CritDamage:        0,
			PhysicalDamage:    0,
			Heal:              0,
		},
		Goblet: map[StatType]StatProfit{
			Attack:            0,
			AttackPercentage:  0,
			Health:            0,
			HealthPercentage:  0,
			Defence:           0,
			DefencePercentage: 0,
			ElementalMastery:  0,
			EnergyRecharge:    0,
			ElementalDamage:   0,
			CritRate:          0,
			CritDamage:        0,
			PhysicalDamage:    0,
			Heal:              0,
		},
		Circlet: map[StatType]StatProfit{
			Attack:            0,
			AttackPercentage:  0,
			Health:            0,
			HealthPercentage:  0,
			Defence:           0,
			DefencePercentage: 0,
			ElementalMastery:  0,
			EnergyRecharge:    0,
			ElementalDamage:   0,
			CritRate:          0,
			CritDamage:        0,
			PhysicalDamage:    0,
			Heal:              0,
		},
	}
}

func (w *CharacterArtifactProfits) Scan(value interface{}) error {
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	err := json.Unmarshal(bytes, w)
	return err
}

func (w *CharacterArtifactProfits) Value() (driver.Value, error) {
	bytes, err := json.Marshal(w)
	return string(bytes), err
}
