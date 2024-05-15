package models

type Book struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  string `json:"price"`
	Status string `json:"status"`
}
