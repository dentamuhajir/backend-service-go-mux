package repository

import (
	"database/sql"

	"github.com/dentamuhajir/backend-service-go-mysql/internal/model"
)

type ArticleRepository struct {
	db *sql.DB
}

func NewArticleRepository(db *sql.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

func (r *ArticleRepository) GetHeadlineArticle() ([]model.Article, error) {
	query := "SELECT id, title, author, category, content_body, photo, photographer, published_date FROM articles ORDER BY RAND() LIMIT 2"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []model.Article
	for rows.Next() {
		var article model.Article
		err := rows.Scan(
			&article.ID,
			&article.Title,
			&article.Author,
			&article.Category,
			&article.ContentBody,
			&article.Photo,
			&article.Photographer,
			&article.PublishedDate,
		)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func (r *ArticleRepository) GetDetailArticle() ([]model.Article, error) {
	query := "SELECT id, title, author, category, content_body, photo, photographer, published_date FROM articles"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []model.Article
	for rows.Next() {
		var article model.Article
		err := rows.Scan(
			&article.ID,
			&article.Title,
			&article.Author,
			&article.Category,
			&article.ContentBody,
			&article.Photo,
			&article.Photographer,
			&article.PublishedDate,
			//&article.IsPublished,
			// &article.CreatedAt,
			// &article.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func (r *ArticleRepository) GetArticleList() ([]model.Article, error) {
	query := "SELECT id, title, author, category, content_body, photo, photographer, published_date FROM articles"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []model.Article
	for rows.Next() {
		var article model.Article
		err := rows.Scan(
			&article.ID,
			&article.Title,
			&article.Author,
			&article.Category,
			&article.ContentBody,
			&article.Photo,
			&article.Photographer,
			&article.PublishedDate,
			//&article.IsPublished,
			// &article.CreatedAt,
			// &article.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}
