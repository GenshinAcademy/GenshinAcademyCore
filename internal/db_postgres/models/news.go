package db_models

import "time"

type News struct {
	Id            DBKey `gorm:"primaryKey"`
	TitleId       DBKey `gorm:"column:title"`
	Title         String
	DescriptionId DBKey `gorm:"column:description"`
	Description   String
	PreviewUrl    string    `gorm:"column:preview_url;type:varchar"`
	RedirectUrl   string    `gorm:"column:redirect_url;type:varchar"`
	CreatedAt     time.Time `gorm:"column:created_at"`
}
