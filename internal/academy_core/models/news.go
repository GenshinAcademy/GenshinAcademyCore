package models

import (
	url "ga/internal/academy_core/value_objects/url"
	"time"
)

type News struct {
	AcademyModel
	Title       string
	Description string
	Preview     string
	RedirectUrl url.Url
	CreatedAt   time.Time
}
