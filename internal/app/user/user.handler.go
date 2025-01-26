package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dentamuhajir/backend-service-go-mysql/internal/app/user"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	userService user.UserService
}

func NewUserHandler(s user.UserService) *UserHandler {
	return &UserHandler{
		userService: s,
	}
}

func (h *UserHandler) RegisterRoute(router *mux.Router) {
	router.HandleFunc("/users", h.getListUsers).Methods("GET")
}

func (h *UserHandler) getListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.GetListUser()
	if err != nil {
		log.Fatalf("error handling User service. Err: %v", err)
	}

	jsonResp, err := json.Marshal(users)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
