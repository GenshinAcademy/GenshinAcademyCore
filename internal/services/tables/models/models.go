package models

import (
	"ga/internal/academy_core/models"
	url "ga/internal/academy_core/value_objects/url"
	"ga/pkg/genshin_core/models/languages"
)

type TablesLocalized struct {
	Id          models.AcademyId              `json:"id,omitempty"`
	Title       map[languages.Language]string `json:"title,omitempty"`
	Description map[languages.Language]string `json:"description,omitempty"`
	Icon        string                        `json:"icon,omitempty"`
	Redirect    url.Url                       `json:"redirect,omitempty"`
}
