package article

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dentamuhajir/backend-service-go-mysql/models"
	"github.com/dentamuhajir/backend-service-go-mysql/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db: db}
}

func (handler *Handler) RegisterRoute(router *mux.Router) {
	router.HandleFunc("/article", handler.getArticle).Methods("GET")
}

func (handler *Handler) getArticle(w http.ResponseWriter, r *http.Request) {
	rows, err := handler.db.Query("SELECT id, title, body, published, category_id FROM articles")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		article.Slug = utils.GenerateSlug(article.Title)

		articles = append(articles, article)
	}

	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(articles); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (handler *Handler) handlerArticle(w http.ResponseWriter, r *http.Request) {
	var article []models.Article
	article = []models.Article{
		{
			ID:         1,
			Title:      "Article 1",
			Body:       "Body 1",
			Published:  true,
			CategoryID: 1,
		}, {
			ID:         2,
			Title:      "Article 2",
			Body:       "Body 2",
			Published:  true,
			CategoryID: 1,
		},
	}

	out, err := json.Marshal(article)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
