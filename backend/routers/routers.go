package routers

import (
	"backend/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	routers := mux.NewRouter()

	routers.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	routers.HandleFunc("/books/{id}", handlers.GetBook).Methods("GET")
	routers.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	return routers
}
