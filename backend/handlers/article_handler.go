package handlers

import (
	"backend/models"
	"encoding/json"
	"net/http"
	"sync"
)

var (
	articles = []models.Article{
		{ID: 1, Title: "New World", Content: "New World is blablabla"},
	}
	mu     sync.Mutex
	nextID = 2
)

func GetArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mu.Lock()
	defer mu.Unlock()
	json.NewEncoder(w).Encode(articles)
}
