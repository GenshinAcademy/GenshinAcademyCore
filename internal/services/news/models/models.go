package models

import (
	"ga/pkg/genshin_core/models/languages"
	// "time"
)

type NewsLocalized struct {
	Title       map[languages.Language]string `json:"title,omitempty" ga:"required,localized" example:"en:Pre-Release,ru:Не прошло и 6 лет..." extensions:"x-order=0"`
	Description map[languages.Language]string `json:"description,omitempty" ga:"required,localized" example:"en:Who needs news when there are such cuties above???,ru:Кому нужны новости когда есть такие милашки???" extensions:"x-order=1"`
	Preview     map[languages.Language]string `json:"preview,omitempty" ga:"required,localized" example:"en:news-en.webp,ru:news-ru.webp" extensions:"x-order=2"`
	Redirect    map[languages.Language]string `json:"redirect,omitempty" ga:"required,localized" example:"en:https://example.com,ru:https://example.com" extensions:"x-order=3"`
	// CreatedAt   time.Time                     `json:"created_at,omitempty" example:"2023-05-21T18:24:28.960356Z" extensions:"x-order=4"`
} //@name NewsLocalized
