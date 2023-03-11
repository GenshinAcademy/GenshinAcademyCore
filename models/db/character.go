package db

type Character struct {
	ID          int `json:"-"`
	CharacterId string
	NameID      int
	Name        Name
	ElementID   int
	Element     Element
	StatsProfit []StatsProfit `gorm:"foreignKey:OwnerID"`
	IconURL     string
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
