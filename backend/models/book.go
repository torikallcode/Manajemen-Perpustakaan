package models

type Book struct {
	Book_id          int    `json:"book_id"`
	Title            string `json:"title"`
	Author           string `json:"author"`
	Publication_year int    `json:"publication_year"`
	Genre            string `json:"genre"`
	Total_copies     int    `json:"total_copies"`
	Isbn             int    `json:"isbn"`
	Language         string `json:"language"`
	Shelf_location   string `json:"shelf_location"`
	Status           string `json:"status"`
	Publisher        string `json:"publisher"`
	Edition          string `json:"edition"`
	Page_count       int    `json:"page_count"`
	Cover_image_url  string `json:"cover_image_url"`
}
