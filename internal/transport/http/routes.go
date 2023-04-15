package http

import (
	"github.com/alibekabdrakhman1/first-lab/internal/transport/http/handler"
	"net/http"

	"github.com/go-chi/chi"
)

func Routes(h *handler.Handler) http.Handler {
	mux := chi.NewRouter()
	mux.Get("/api/users/", h.User.GetAll)
	mux.Get("/api/users/{login}", h.User.Get)
	mux.Delete("/api/users/{login}", h.User.Delete)
	mux.Post("/api/users/create", h.User.Create)
	mux.Post("/api/users/update", h.User.Update)

	mux.Get("/api/books/", h.Book.GetAll)
	mux.Get("/api/books/getByAuthor/{author}", h.Book.GetByAuthor)
	mux.Get("/api/books/getByName/{name}", h.Book.GetByName)
	mux.Get("/api/books/getAvailable", h.Book.GetAvailableBooks)
	mux.Post("/api/books/create", h.Book.Create)
	mux.Post("/api/books/update", h.Book.Update)
	mux.Delete("/api/books/{name}", h.Book.Delete)

	mux.Get("/api/orders/", h.Order.GetAll)
	mux.Get("/api/orders/notReturned", h.Order.GetNotReturned)
	mux.Get("/api/orders/lastMonth", h.Order.GetCountOrders)
	mux.Post("/api/orders/create", h.Order.Create)
	mux.Delete("/api/orders/{orderId}", h.Order.Delete)
	mux.Post("/api/orders/return/{orderId}", h.Order.ReturnBook)

	return mux
}
