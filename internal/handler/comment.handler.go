package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/dentamuhajir/backend-service-go-mysql/internal/service"
	"github.com/gorilla/mux"
)

type CommentHandler struct {
	commentService service.CommentService
}

func NewCommentHandler(s service.CommentService) *CommentHandler {
	return &CommentHandler{commentService: s}
}

func (h *CommentHandler) RegisterRoute(router *mux.Router) {
	router.HandleFunc("/comment/article/{articleID}", h.getCommentByArticleID).Methods("GET")
}

func (h *CommentHandler) getCommentByArticleID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	articleId, _ := strconv.ParseInt(params["articleID"], 10, 64)

	comments, err := h.commentService.GetCommentByArticleID(articleId)

	if err != nil {
		log.Fatalf("error handling Article service. Err: %v", err)
	}

	error := json.NewEncoder(w).Encode(comments)

	if error != nil {
		log.Fatalf("error handling JSON Encode. Err: %v", err)
	}
}
