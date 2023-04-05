package transport

import (
	"net/http"

	"github.com/go-chi/chi"
)

func Routes() http.Handler {
	mux := chi.NewRouter() // роутер инициализируется отдельно
	h := NewHandler() // хендлер должен заходить снаружи
	mux.Get("/api/users", h.User.GetAll)
	mux.Get("/api/users/{id}", h.User.Get)
	
	return mux	

}