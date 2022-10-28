package entities

type Game struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Genre string `json:"genre"`
}
