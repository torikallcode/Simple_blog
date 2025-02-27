package main

import (
	"backend/database"
	"backend/routers"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {

	database.InitDatabase()
	defer database.DB.Close()

	router := routers.SetupRouter()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		Debug:          true,
	})

	handler := c.Handler(router)

	log.Println("Server sedang berjalan di port :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
