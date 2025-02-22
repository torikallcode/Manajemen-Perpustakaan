package handlers

import (
	"backend/databases"
	"backend/models"
	"encoding/json"
	"net/http"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var books []models.Book

	query := "SELECT * FROM books"
	rows, err := databases.DB.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.Book_id, &book.Title, &book.Author, &book.Publication_year, &book.Genre, &book.Total_copies); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		books = append(books, book)
	}

	json.NewEncoder(w).Encode(books)
}
