package models

import (
	"ga/pkg/genshin_core/models/languages"
)

type TablesLocalized struct {
	Title       map[languages.Language]string `json:"title,omitempty" ga:"required,localized" example:"en:Monsters' resistances and shields,ru:Сопротивления и щиты монстров" extensions:"x-order=1"`
	Description map[languages.Language]string `json:"description,omitempty" ga:"required,localized" example:"en:Elemental resistances and gauges of their elemental shields or structures along with notes on their specific mechanics that change those values.,ru:Лиза" extensions:"x-order=2"`
	Icon        string                        `json:"icon,omitempty" ga:"required" example:"shield.webp" extensions:"x-order=3"`
	Redirect    map[languages.Language]string `json:"redirect,omitempty" ga:"required,localized" example:"en:https://example.com,ru:https://example.com" extensions:"x-order=4"`
} //@name TablesLocalized
