package models

import (
	"ga/internal/academy_core/models"
	"ga/pkg/genshin_core/models/languages"
)

type TablesLocalized struct {
	Id          models.AcademyId              `json:"id,omitempty"`
	Title       map[languages.Language]string `json:"title,omitempty" ga:"required,localized"`
	Description map[languages.Language]string `json:"description,omitempty" ga:"required,localized"`
	Icon        string                        `json:"icon,omitempty" ga:"required"`
	Redirect    map[languages.Language]string `json:"redirect,omitempty" ga:"required,localized"`
}
