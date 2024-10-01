package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/dentamuhajir/backend-service-go-mysql/internal/handler"
	"github.com/dentamuhajir/backend-service-go-mysql/internal/repository"
	"github.com/dentamuhajir/backend-service-go-mysql/internal/service"
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

	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userRepository 	:= repository.NewUserRepository(s.db)
	userService 	  := service.NewUserService(*userRepository)
	userHandler 	  := handler.NewUserHandler(*userService)
	userHandler.RegisterRoute(subrouter)

	log.Println("Starting API server", s.addr)
	return http.ListenAndServe(s.addr, router)
}
