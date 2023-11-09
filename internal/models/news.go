package models

import (
	"ga/internal/types"
	"time"
)

type News struct {
	Id          types.NewsId
	Title       string
	Description string
	Preview     string
	RedirectUrl string
	CreatedAt   time.Time
}

type NewsMultilingual struct {
	Id          types.NewsId
	Title       types.LocalizedString
	Description types.LocalizedString
	Preview     types.LocalizedString
	RedirectUrl types.LocalizedString
	CreatedAt   time.Time
}
