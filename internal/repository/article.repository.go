package repository

import (
	"database/sql"
	"log"

	"github.com/dentamuhajir/backend-service-go-mysql/internal/model"
)

type ArticleRepository struct {
	db *sql.DB
}

func NewArticleRepository(db *sql.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

func (r *ArticleRepository) GetHeadlineArticle() ([]model.Headline, error) {
	query := "SELECT id, title, category, photo FROM articles ORDER BY RAND() LIMIT 2"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []model.Headline
	for rows.Next() {
		var article model.Headline
		err := rows.Scan(
			&article.ID,
			&article.Title,
			&article.Category,
			&article.Photo,
		)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func (r *ArticleRepository) GetDetailArticle(id int64) ([]model.Article, error) {
	query := "SELECT id, title, author, category, content_body, photo, photographer, published_date FROM articles WHERE id = ?"
	rows, err := r.db.Query(query, id)
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

func (r *ArticleRepository) SaveArticle(article model.StoreArticle) (message string) {

	stmt, err := r.db.Prepare("INSERT INTO articles (id, title, author, category, content_body, photo, photographer, is_published) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(10001, article.Title, article.Author, article.Category, article.ContentBody, article.Photo, article.Photographer, true)

	if err != nil {
		log.Fatal(err)
		return "Something bad happened on the server "
	}

	return "Success insert data"
}

func (r *ArticleRepository) GetArticleList() ([]model.ListArticle, error) {
	query := "SELECT id, title, category, photo FROM articles"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []model.ListArticle
	for rows.Next() {
		var article model.ListArticle
		err := rows.Scan(
			&article.ID,
			&article.Title,
			&article.Category,
			&article.Photo,
		)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}
