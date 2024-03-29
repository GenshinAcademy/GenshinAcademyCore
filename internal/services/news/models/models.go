package models

import (
	"ga/internal/academy_core/models"
	"ga/pkg/genshin_core/models/languages"
	"time"
)

type NewsLocalized struct {
	Id          models.AcademyId              `json:"id,omitempty"`
	Title       map[languages.Language]string `json:"title,omitempty" ga:"required,localized"`
	Description map[languages.Language]string `json:"description,omitempty" ga:"required,localized"`
	Preview     map[languages.Language]string `json:"preview,omitempty" ga:"required,localized"`
	Redirect    map[languages.Language]string `json:"redirect,omitempty" ga:"required,localized"`
	CreatedAt   time.Time                     `json:"created_at,omitempty"`
}
