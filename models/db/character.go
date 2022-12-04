package db

type Character struct {
	ID          int `json:"-"`
	CharacterId string
	Name        Name          `gorm:"foreignKey:ID"`
	Element     Element       `gorm:"many2many:character_elements"`
	StatsProfit []StatsProfit `gorm:"foreignKey:OwnerID"`
}

type Name struct {
	ID      int `json:"-"`
	English string
	Russian string
}

type Element struct {
	ID   int `json:"-"`
	Name string
}
