package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dentamuhajir/backend-service-go-mysql/internal/service"
	"github.com/gorilla/mux"
)

type ArticleHandler struct {
	articleService service.ArticleService
}

func NewArticleHandler(s service.ArticleService) *ArticleHandler {
	return &ArticleHandler{
		articleService: s,
	}
}

func (h *ArticleHandler) RegisterRoute(router *mux.Router) {
	router.HandleFunc("/article", h.getListArticle).Methods("GET")
	router.HandleFunc("/article/headline", h.getHeadlineArticle).Methods("GET")
}

func (h *ArticleHandler) getListArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	articles, err := h.articleService.GetListArticle()
	if err != nil {
		log.Fatalf("error handling Article service. Err: %v", err)
	}

	error := json.NewEncoder(w).Encode(articles)

	if error != nil {
		log.Fatalf("error handling JSON Encode. Err: %v", err)
	}
}

func (h *ArticleHandler) getHeadlineArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	articles, err := h.articleService.GetHeadlineArticle()
	if err != nil {
		log.Fatalf("error handling headline article. Err: %v", err)
	}

	error := json.NewEncoder(w).Encode(articles)
	if error != nil {
		log.Fatalf("error handling JSON Encode. Err: %v", err)
	}
}
