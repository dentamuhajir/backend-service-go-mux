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

func NewArticleHandler(s service.ArticleService) *ArticleHandler{
	return &ArticleHandler{
		articleService: s,
	}
}

func (h *ArticleHandler) RegisterRoute(router *mux.Router) {
	router.HandleFunc("/article", h.getListArticle).Methods("GET")
}

func (h *ArticleHandler) getListArticle(w http.ResponseWriter, r *http.Request) {
	articles, err := h.articleService.GetListArticle();
	if err != nil {
		log.Fatalf("error handling Article service. Err: %v", err)
	}

	jsonResp, err := json.Marshal(articles)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}