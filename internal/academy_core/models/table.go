package models

import (
	url "ga/internal/academy_core/value_objects/url"
)

type Table struct {
	AcademyModel
	Title       string
	Description string
	Icon        string
	RedirectUrl url.Url
}
