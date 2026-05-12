package auth

import (
	"net/http"

	"github.com/go-chi/render"
)

type Handler struct {
	// authService *service.AuthService  // later
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, map[string]string{
		"message": "Register endpoint - to be implemented",
	})
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, map[string]string{
		"message": "Login endpoint - to be implemented",
	})
}
