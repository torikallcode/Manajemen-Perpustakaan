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
	defer rows.Close() // Pastikan rows ditutup setelah selesai

	for rows.Next() {
		var book models.Book
		// Gunakan sql.NullString untuk menangani NULL
		var language sql.NullString

		if err := rows.Scan(
			&book.Book_id,
			&book.Title,
			&book.Author,
			&book.Publication_year,
			&book.Genre,
			&book.Total_copies,
			&book.Isbn,
			&language, // Gunakan variable language yang baru
			&book.Shelf_location,
			&book.Status,
			&book.Publisher,
			&book.Edition,
			&book.Page_count,
			&book.Cover_image_url,
		); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Hanya set Language jika nilainya valid
		if language.Valid {
			book.Language = language.String
		}

		books = append(books, book)
	}

	// Tambahkan pengecekan error setelah iterasi
	if err = rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	query := "UPDATE books SET title = ?, author = ?, publication_year = ?, genre = ?, total_copies = ?, isbn = ?, language = ?, shelf_location = ?, status = ?, publisher = ?, edition = ?, page_count = ?, cover_image_url WHERE book_id = ?"
	_, err = databases.DB.Exec(query, book.Title, book.Author, book.Publication_year, book.Genre, book.Total_copies, book.Isbn, book.Language, book.Shelf_location, book.Status, book.Publisher, book.Edition, book.Page_count, book.Cover_image_url, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Fetch the updated book to ensure consistency
	var updatedBook models.Book
	err = databases.DB.QueryRow("SELECT * FROM books WHERE book_id = ?", id).Scan(&updatedBook.Book_id, &updatedBook.Title, &updatedBook.Author, &updatedBook.Publication_year, &updatedBook.Genre, &updatedBook.Total_copies)
	if err != nil {
		http.Error(w, "Failed to fetch updated book", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedBook)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	query := "DELETE FROM books WHERE book_id = ?"
	result, err := databases.DB.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "rows not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
