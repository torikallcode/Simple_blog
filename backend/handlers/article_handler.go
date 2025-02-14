package handlers

import (
	"backend/database"
	"backend/models"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

var (
	mu sync.Mutex
)

func GetArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mu.Lock()
	defer mu.Unlock()

	var articles []models.Article

	query := "SELECT id, title, content FROM article"
	rows, err := database.DB.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var article models.Article
		if err := rows.Scan(&article.ID, &article.Title, article.Content); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		articles = append(articles, article)
	}

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

	var article models.Article

	query := "SELECT id , title, content FROM article WHERE id =?"
	err = database.DB.QueryRow(query, id).Scan(&article.ID, &article.Title, &article.Content)
	if err == sql.ErrNoRows {
		http.Error(w, "rown not founf", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(article)
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

	article.ID = len(articles) + 1
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

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
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
			articles = append(articles[:index], articles[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "article not found", http.StatusNotFound)
}
