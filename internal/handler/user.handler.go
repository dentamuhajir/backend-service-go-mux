package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dentamuhajir/backend-service-go-mysql/internal/service"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	userService service.UserService 
}

func NewUserHandler(s service.UserService) *UserHandler{
	return &UserHandler{
		userService: s,
	}
}

func (h *UserHandler) RegisterRoute(router *mux.Router) {
	router.HandleFunc("/users", h.getListUsers).Methods("GET")
}

func (h *UserHandler) getListUsers(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}