package api

import (
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
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
	log.Println("Starting API server", s.addr)
	return http.ListenAndServe(s.addr, router)
}
