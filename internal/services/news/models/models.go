package models

import (
	"ga/internal/academy_core/models"
	url "ga/internal/academy_core/value_objects/url"
	"ga/pkg/genshin_core/models/languages"
	"time"
)

type NewsLocalized struct {
	Id          models.AcademyId              `json:"id,omitempty"`
	Title       map[languages.Language]string `json:"title,omitempty"`
	Description map[languages.Language]string `json:"description,omitempty"`
	Preview     string                        `json:"preview,omitempty"`
	Redirect    url.Url                       `json:"redirect,omitempty"`
	CreatedAt   time.Time                     `json:"created_at,omitempty"`
}
