package entities

type Game struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Developer string `json:"developer"`
	Started   bool   `json:"started"`
	Finished  bool   `json:"finished"`
}
