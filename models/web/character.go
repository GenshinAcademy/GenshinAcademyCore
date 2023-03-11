package web

type Character struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Element string `json:"element"`
	IconURL string `json:"icon_url"`
}
