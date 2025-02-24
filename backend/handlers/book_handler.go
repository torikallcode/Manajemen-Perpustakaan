package handlers

import (
	"backend/databases"
	"backend/models"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var book models.Book

	query := "SELECT book_id, title, author, publication_year, genre, total_copies FROM books WHERE book_id = ?"
	err = databases.DB.QueryRow(query, id).Scan(&book.Book_id, &book.Title, &book.Author, &book.Publication_year, &book.Genre, &book.Total_copies)
	if err == sql.ErrNoRows {
		http.Error(w, "rows not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "invalid book", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO books (title, author, publication_year, genre, total_copies) VALUE (?,?,?,?,?)"
	result, err := databases.DB.Exec(query, &book.Title, &book.Author, &book.Publication_year, &book.Genre, &book.Total_copies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	book.Book_id = int(id)
	json.NewEncoder(w).Encode(book)
}
