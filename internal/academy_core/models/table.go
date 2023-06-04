package models

type Table struct {
	AcademyModel
	Title       string `example:"Monsters' resistances and shields" extensions:"x-order=2"`
	Description string `example:"Elemental resistances and gauges of their elemental shields or structures along with notes on their specific mechanics that change those values." extensions:"x-order=3"`
	Icon        string `example:"https://example.com" extensions:"x-order=4"`
	RedirectUrl string `example:"https://example.com" extensions:"x-order=5"`
} //@name Table
