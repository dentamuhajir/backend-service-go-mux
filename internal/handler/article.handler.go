package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dentamuhajir/backend-service-go-mysql/internal/service"
	"github.com/gorilla/mux"
	"github.com/brianvoe/gofakeit/v7"
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
	router.HandleFunc("/article/categories", h.getArticleByCategory).Methods("GET")
	router.HandleFunc("/article/detail/{id}/{slug}", h.getDetailArticle).Methods("GET")
	router.HandleFunc("/article/dummy-data", h.generateDummyData).Methods("GET")
}

func (h *ArticleHandler) getDetailArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	slug := params["slug"]
	fmt.Fprintln(w, "id: ", id)
	fmt.Fprintln(w, "slug: ", slug)
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
		log.Fatalf("Error handling JSON Encode. Err: %v", err)
	}
}

func (h *ArticleHandler) getArticleByCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	articles, err := h.articleService.GetListArticleGroupByCategory()
	if err != nil {
		log.Fatalf("Error handling Article by categories. Err: %v", err)
	}
	error := json.NewEncoder(w).Encode(articles)
	if error != nil {
		log.Fatalf("Error handling JSON Encode: %v", err)
	}
}

func (h *ArticleHandler) generateDummyData(w http.ResponseWriter, r *http.Request) {
	name := gofakeit.Name()
	email := gofakeit.Email()          
	phone := gofakeit.Phone()
	company := gofakeit.Company()  
	
	fmt.Fprintln(w, name , email, phone, company)
}
