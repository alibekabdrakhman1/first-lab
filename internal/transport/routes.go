package transport

import (
	"net/http"

	"github.com/go-chi/chi"
)

func Routes() http.Handler {
	mux := chi.NewRouter()
	h := NewHandler()
	mux.Get("/api/users", h.User.GetAll)
	mux.Get("/api/users/{id}", h.User.Get)
	
	return mux	

}