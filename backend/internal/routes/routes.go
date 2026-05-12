package routes

import (
	"backend/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func Setup(r *chi.Mux, h *handler.Handler) {
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, map[string]string{"status": "healthy"})
	})

	// ------------v1------------

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/register", h.Auth.Register)
		r.Post("/login", h.Auth.Login)
	})

	r.Group(func(r chi.Router) {
		// r.Use(middleware.Auth())

		r.Route("/users", func(r chi.Router) {
			// r.Get("/", h.User.GetAll)
			// r.Get("/{id}", h.User.GetByID)
			// r.Post("/", h.User.Create)
			// r.Put("/{id}", h.User.Update)
			// r.Delete("/{id}", h.User.Delete)
		})
	})
}
