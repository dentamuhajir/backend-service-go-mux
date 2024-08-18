package article

import (
	"encoding/json"
	"fmt"
	"github.com/dentamuhajir/backend-service-go-mysql/models"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (handler *Handler) RegisterRoute(router *mux.Router) {
	router.HandleFunc("/article", handler.handlerArticle).Methods("GET")

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
