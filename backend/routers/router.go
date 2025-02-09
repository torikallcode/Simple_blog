package routers

import (
	"backend/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	routers := mux.NewRouter()

	routers.HandleFunc("/articles", handlers.GetArticles).Methods("GET")
	routers.HandleFunc("/articles/{id}", handlers.GetArticle).Methods("GET")
	routers.HandleFunc("/articles", handlers.CreateArticle).Methods("POST")
	routers.HandleFunc("/articles/{id}", handlers.UpdateArticle).Methods("PUT")
	routers.HandleFunc("/articles/{id}", handlers.DeleteArticle).Methods("DELETE")

	return routers
}
