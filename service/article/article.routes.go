package article

import (
	"database/sql"
	"encoding/json"
	"log"
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
	router.HandleFunc("/article/create", handler.postArticle).Methods("GET")
}

func (handler *Handler) postArticle(w http.ResponseWriter, r *http.Request) {

	requestBody := models.Article{
		ID:         90,
		Title:      "Testing insert data",
		Body:       "Testing insert data",
		Published:  true,
		CategoryID: 2,
		Slug:       "testing-insert-data",
	}

	message, err := handler.service.postArticle(requestBody)

	if err != nil {
		log.Println(err)
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(message); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (handler *Handler) getArticle(w http.ResponseWriter, r *http.Request) {
	var articles []models.Article = handler.service.getArticle()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(articles); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
