package repository

import (
	"backend/models"
	"context"
	"database/sql"
)

type ArticleRepository interface {
	GetArticles(ctx context.Context) ([]models.Article, error)
	GetArticle(ctx context.Context, id int) (models.Article, error)
	CreateArticle(ctx context.Context, article models.Article) (int, error)
	UpdateArticle(ctx context.Context, id int, article models.Article) error
	DeleteArticle(ctx context.Context, id int) error
}

type articleRepository struct {
	db *sql.DB
}

func NewArticleRepository(db *sql.DB) ArticleRepository {
	return &articleRepository{db: db}
}

func (r *articleRepository) GetArticles(ctx context.Context) ([]models.Article, error) {
	query := "SELECT id, title, content FROM article"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []models.Article
	for rows.Next() {
		var article models.Article
		if err := rows.Scan(&article.ID, &article.Title, &article.Content); err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func (r *articleRepository) GetArticle(ctx context.Context, id int) (models.Article, error) {
	query := "SELECT id, title, content FROM article WHERE id = ?"
	var article models.Article
	err := r.db.QueryRowContext(ctx, query, id).Scan(&article.ID, &article.Title, &article.Content)
	if err == sql.ErrNoRows {
		return models.Article{}, err
	}
	return article, err
}

func (r *articleRepository) CreateArticle(ctx context.Context, article models.Article) (int, error) {
	query := "INSERT INTO article (title, content) VALUES (?, ?)"
	result, err := r.db.ExecContext(ctx, query, article.Title, article.Content)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *articleRepository) UpdateArticle(ctx context.Context, id int, article models.Article) error {
	query := "UPDATE article SET title = ?, content = ? WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, article.Title, article.Content, id)
	return err
}

func (r *articleRepository) DeleteArticle(ctx context.Context, id int) error {
	query := "DELETE FROM article WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
