package expenses

type expense struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Amount float32 `json:"amount"`
	NOTE   string  `json:"note"`
	Tags   string  `json:"tags"`
}
