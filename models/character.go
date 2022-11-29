package models

type Character struct {
	ID          int    `json:"id" gorm:"primary_key"`
	CharacterId string `json:"character_id"`
	Name        string `json:"name"`
	Element     string `json:"element"`
	IconUrl     string `json:"icon_url"`
}
