package characters

import common "genshinacademycore"

type CharacterStats struct {
	common.Entity
	Attack         int32          `json:"attack"`
	Health         int32          `json:"health"`
	Defense        float32        `json:"defense"`
	CriticalRate   float32        `json:"critRate"`
	CriticalDamage float32        `json:"critDamage"`
	Bonus          CharacterBonus `json:"bonus"`
}
