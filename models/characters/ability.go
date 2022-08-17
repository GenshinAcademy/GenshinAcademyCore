package characters

import common "genshinacademycore"

const (
	MaximimAbilityLevel uint8 = 15
	MinimumAbilityLevel uint  = 1
)

type Skill struct {
	common.Entity
	Cooldown    float32 `json:"cooldown"`
	Description string  `json:"description"`
	LoreText    string  `json:"loreText"`
}
