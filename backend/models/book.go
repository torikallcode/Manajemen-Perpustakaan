package models

type Book struct {
	Book_id          int    `json:"book_id"`
	Title            string `json:"title"`
	Author           string `json:"author"`
	Publication_year int    `json:"publication_year"`
	Genre            string `json:"genre"`
	Total_copies     int    `json:"total_copies"`
}
