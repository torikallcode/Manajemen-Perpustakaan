package routers

import (
	"backend/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	routers := mux.NewRouter()

	routers.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	return routers
}
