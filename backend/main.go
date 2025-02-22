package main

import (
	"backend/databases"
	"backend/routers"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	databases.InitDatabase()
	defer databases.DB.Close()

	routers := routers.SetupRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http:localhost:5173"},
		AllowedMethods:   []string{"GET"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		Debug:            true,
	})

	handler := c.Handler(routers)

	log.Println("Server sedang berjalan di port:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))

}
