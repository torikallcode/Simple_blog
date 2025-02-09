package handlers

import (
	"backend/models"
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
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

func GetArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	mu.Lock()
	defer mu.Unlock()
	if err != nil {
		http.Error(w, "invalid article", http.StatusBadRequest)
		return
	}
	for _, item := range articles {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "article not found", http.StatusNotFound)
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var article models.Article
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	article.ID = nextID
	articles = append(articles, article)
	json.NewEncoder(w).Encode(article)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "invalid article", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for index, item := range articles {
		if item.ID == id {
			var UpdateArticle models.Article
			if err := json.NewDecoder(r.Body).Decode(&UpdateArticle); err != nil {
				http.Error(w, "invalid article", http.StatusBadRequest)
				return
			}
			UpdateArticle.ID = id
			articles[index] = UpdateArticle
			json.NewEncoder(w).Encode(UpdateArticle)
			return
		}
	}
	http.Error(w, "article not found", http.StatusNotFound)
}
