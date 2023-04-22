package db_models

import "time"

type News struct {
	Id            DBKey `gorm:"primaryKey"`
	TitleId       DBKey `gorm:"column:title"`
	Title         String
	DescriptionId DBKey `gorm:"column:description"`
	Description   String
	PreviewId     DBKey `gorm:"column:preview_id"`
	PreviewUrl    String
	RedirectId    DBKey `gorm:"column:redirect_id"`
	RedirectUrl   String
	CreatedAt     time.Time `gorm:"column:created_at"`
}
