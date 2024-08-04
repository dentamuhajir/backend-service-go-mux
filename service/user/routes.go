package user

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.HandlerLogin).Methods("GET")
	router.HandleFunc("/register", h.HandlerRegister).Methods("POST")
}

func (h *Handler) HandlerLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login Route Called")
}

func (h *Handler) HandlerRegister(w http.ResponseWriter, r *http.Request) {

}
