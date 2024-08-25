package article

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/dentamuhajir/backend-service-go-mysql/models"
	"github.com/gorilla/mux"
)

type Handler struct {
	service ArticleService
	db      *sql.DB
}

func NewHandler(service ArticleService, db *sql.DB) *Handler {
	return &Handler{service: service, db: db}
}

func (handler *Handler) RegisterRoute(router *mux.Router) {
	router.HandleFunc("/article", handler.getArticle).Methods("GET")
}

func (handler *Handler) getArticle(w http.ResponseWriter, r *http.Request) {
	var articles []models.Article = handler.service.getArticle()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(articles); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// func (handler *Handler) handlerArticle(w http.ResponseWriter, r *http.Request) {
// 	var article []models.Article
// 	article = []models.Article{
// 		{
// 			ID:         1,
// 			Title:      "Article 1",
// 			Body:       "Body 1",
// 			Published:  true,
// 			CategoryID: 1,
// 		}, {
// 			ID:         2,
// 			Title:      "Article 2",
// 			Body:       "Body 2",
// 			Published:  true,
// 			CategoryID: 1,
// 		},
// 	}

// 	out, err := json.Marshal(article)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(out)
// }
