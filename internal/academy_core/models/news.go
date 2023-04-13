package models

import (
	"time"
)

type News struct {
	AcademyModel
	Title       string
	Description string
	Preview     string
	RedirectUrl string
	CreatedAt   time.Time
}
