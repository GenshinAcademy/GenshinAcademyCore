package models

import (
	"ga/internal/types"
)

type Table struct {
	Id          types.TableId
	Title       string
	Description string
	IconUrl     string
	RedirectUrl string
}

type TableMultilingual struct {
	Id          types.TableId
	Title       types.LocalizedString
	Description types.LocalizedString
	IconUrl     string
	RedirectUrl types.LocalizedString
}
