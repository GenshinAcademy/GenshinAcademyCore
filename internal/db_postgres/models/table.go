package db_models

type Table struct {
	Id            DBKey `gorm:"primaryKey"`
	TitleId       DBKey `gorm:"column:title"`
	Title         String
	DescriptionId DBKey `gorm:"column:description"`
	Description   String
	PreviewUrl    string    `gorm:"column:preview_url;type:varchar"`
	RedirectUrl   string    `gorm:"column:redirect_url;type:varchar"`
}