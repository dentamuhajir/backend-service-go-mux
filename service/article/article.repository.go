package article

import (
	"database/sql"

	"github.com/dentamuhajir/backend-service-go-mysql/models"
	"github.com/dentamuhajir/backend-service-go-mysql/utils"
)

type ArticleRepository struct {
	db *sql.DB
}

func NewArticleRepository(db *sql.DB) *ArticleRepository {
	return &ArticleRepository{db}
}

func (r ArticleRepository) insertArticle(requestBody models.Article) string {
	stmt, err := r.db.Prepare("INSERT INTO articles(id, title, body, published, category_id) VALUES (?,?,?,?,?)")

	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	stmt.Exec(
		requestBody.ID,
		requestBody.Title,
		requestBody.Body,
		requestBody.Published,
		requestBody.CategoryID,
	)

	return "Insert Article Successfull"
}

func (r ArticleRepository) getArticle() ([]models.Article, error) {
	rows, err := r.db.Query("SELECT id, title, body, published, category_id FROM articles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []models.Article
	for rows.Next() {
		var article models.Article
		// reading for the row of queryp
		if err := rows.Scan(
			&article.ID,
			&article.Title,
			&article.Body,
			&article.Published,
			&article.CategoryID,
		); err != nil {
			return nil, err
		}

		article.Slug = utils.GenerateSlug(article.Title)

		articles = append(articles, article)
	}

	return articles, err
}
