package models

import (
	"time"
)

type News struct {
	AcademyModel
	Title       string    `example:"Pre-Release" extensions:"x-order=1"`
	Description string    `example:"Who needs news when there are such cuties above???" extensions:"x-order=2"`
	Preview     string    `example:"https://example.com" extensions:"x-order=3"`
	RedirectUrl string    `example:"https://example.com" extensions:"x-order=4"`
	CreatedAt   time.Time `example:"2023-05-21T18:24:28.960356Z" extensions:"x-order=5"`
} //@name News
