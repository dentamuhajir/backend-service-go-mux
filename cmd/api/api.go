package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

// constructor function
func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Start() error {
	router := mux.NewRouter()
	// subrouter := router.PathPrefix("/api/v1").Subrouter()
	// userHandler := user.NewHandler()

	// articleRepository := article.NewArticleRepository(s.db)
	// articleService := article.NewArticleService(*articleRepository)
	// articleHandler := article.NewHandler(*articleService, s.db)
	// articleHandler.RegisterRoute(subrouter)
	// userHandler.RegisterRoutes(subrouter)
	log.Println("Starting API server", s.addr)
	return http.ListenAndServe(s.addr, router)
}
