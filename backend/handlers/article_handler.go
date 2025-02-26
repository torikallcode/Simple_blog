package handlers

import (
	"backend/models"
	"backend/repository"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type ArticleHandler struct {
	repo repository.ArticleRepository
}

func NewArticleHandler(repo repository.ArticleRepository) *ArticleHandler {
	return &ArticleHandler{repo: repo}
}

func (h *ArticleHandler) GetArticles(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	articles, err := h.repo.GetArticles(ctx)
	if err != nil {
		log.Printf("Error fetching articles: %v", err)
		http.Error(w, "Failed to fetch articles", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

func (h *ArticleHandler) GetArticle(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}

	article, err := h.repo.GetArticle(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Article not found", http.StatusNotFound)
		} else {
			log.Printf("Error fetching article: %v", err)
			http.Error(w, "Failed to fetch article", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}

func (h *ArticleHandler) CreateArticle(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	var article models.Article
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	id, err := h.repo.CreateArticle(ctx, article)
	if err != nil {
		log.Printf("Error creating article: %v", err)
		http.Error(w, "Failed to create article", http.StatusInternalServerError)
		return
	}

	article.ID = id
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(article)
}

func (h *ArticleHandler) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}

	var article models.Article
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.repo.UpdateArticle(ctx, id, article); err != nil {
		log.Printf("Error updating article: %v", err)
		http.Error(w, "Failed to update article", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}

func (h *ArticleHandler) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}

	if err := h.repo.DeleteArticle(ctx, id); err != nil {
		log.Printf("Error deleting article: %v", err)
		http.Error(w, "Failed to delete article", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
