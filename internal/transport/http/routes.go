package http

import (
	"net/http"

	"github.com/go-chi/chi"
)

func Routes(h *Handler) http.Handler {
	mux := chi.NewRouter()
	mux.Get("/api/users", h.User.GetAll)
	mux.Get("/api/users/{login}", h.User.Get)
	mux.Delete("/api/users/{login}", h.User.Delete)
	mux.Post("/api/users/create", h.User.Create)
	mux.Post("/api/users/update", h.User.Update)
	return mux
}
