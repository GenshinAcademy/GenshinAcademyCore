package models

type Character struct {
	ID      int    `json:"id" gorm:"primary_key"`
	Name    string `json:"name"`
	Element string `json:"element"`
	IconUrl string `json:"icon_url"`
}
