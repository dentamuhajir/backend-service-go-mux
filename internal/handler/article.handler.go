package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/dentamuhajir/backend-service-go-mysql/internal/service"
	"github.com/gorilla/mux"
)

type ArticleHandler struct {
	articleService service.ArticleService
	GenericResponse
}

type GenericResponse struct {
	Status string `json:"status"`
	Code   int16  `json:"code"`
	Data   any    `json:"data"`
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
	// Convert string to int64
	id, _ := strconv.ParseInt(params["id"], 10, 64)
	articles, err := h.articleService.GetDetailArticle(id)

	if err != nil {
		log.Fatalf("error handling Article service. Err: %v", err)
	}

	error := json.NewEncoder(w).Encode(articles)

	if error != nil {
		log.Fatalf("error handling JSON Encode. Err: %v", err)
	}

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

	h.GenericResponse.Status = "Success"
	h.GenericResponse.Code = 200
	h.GenericResponse.Data = articles

	error := json.NewEncoder(w).Encode(h.GenericResponse)
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

	h.GenericResponse.Status = "Success"
	h.GenericResponse.Code = 200
	h.GenericResponse.Data = articles

	error := json.NewEncoder(w).Encode(h.GenericResponse)
	if error != nil {
		log.Fatalf("Error handling JSON Encode: %v", err)
	}
}

func (h *ArticleHandler) generateDummyData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	message := h.articleService.GenerateDummyData()
	fmt.Fprintln(w, "message: ", message)
}
