package handlers

import (
	"backend/database"
	"backend/models"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetArticles(w http.ResponseWriter, r *http.Request) {
	// Set header response sebagai JSON
	// Eksekusi query SELECT untuk mengambil semua data
	// Membaca hasil query baris per baris
	// Menyimpan setiap baris ke dalam slice lists
	// Mengirim seluruh data lists sebagai JSON response
	w.Header().Set("Content-Type", "application/json")
	var articles []models.Article

	query := "SELECT id, title, content FROM article"
	rows, err := database.DB.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for rows.Next() {
		var article models.Article
		if err := rows.Scan(&article.ID, &article.Title, &article.Content); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		articles = append(articles, article)
	}
	json.NewEncoder(w).Encode(articles)
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	// 	Set header response sebagai JSON
	// Ambil parameter ID dari URL
	// Konversi ID dari string ke integer
	// Eksekusi query SELECT untuk mengambil data spesifik
	// Memindahkan data ke variabel list
	// Mengirim data list sebagai JSON response
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "invalid article", http.StatusBadRequest)
		return
	}

	var article models.Article
	query := "SELECT id, title, content FROM article WHERE id = ?"
	err = database.DB.QueryRow(query, id).Scan(&article.ID, &article.Title, &article.Content)
	if err == sql.ErrNoRows {
		http.Error(w, "rows not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)

}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	// Set header response sebagai JSON
	// Decode data JSON dari request body
	// Validasi input
	// Eksekusi query INSERT ke database
	// Ambil ID dari baris yang baru dimasukkan
	// Kirim data list yang baru dibuat sebagai response
	w.Header().Set("Content-Type", "application/json")
	var article models.Article
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO article (title, content) VALUE (?,?)"
	result, err := database.DB.Exec(query, &article.Title, &article.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	article.ID = int(id)
	json.NewEncoder(w).Encode(article)

}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	// Set header response sebagai JSON
	// Ambil parameter ID dari URL
	// Konversi ID dari string ke integer
	// Decode data JSON dari request body
	// Validasi input
	// Eksekusi query UPDATE ke database
	// Kirim data list yang diupdate sebagai response
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "invalid request", http.StatusInternalServerError)
		return
	}

	var article models.Article
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	query := "UPDATE article SET title = ?, content = ? WHERE id = ?"
	_, err = database.DB.Exec(query, &article.Title, &article.Content, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)

}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	// Set header response sebagai JSON
	// Ambil parameter ID dari URL
	// Konversi ID dari string ke integer
	// Eksekusi query DELETE ke database
	// Periksa jumlah baris yang terpengaruh
	// Kirim response sesuai hasil operasi
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	query := "DELETE FROM article WHERE id = ?"
	result, err := database.DB.Exec(query, id)
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
