package handler

import (
	"encoding/json"
	"github.com/alibekabdrakhman1/first-lab/internal/model"
	"github.com/alibekabdrakhman1/first-lab/internal/service"
	"net/http"
	"strings"
)

type IBookHandler interface {
	GetByAuthor(http.ResponseWriter, *http.Request)
	GetByName(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	GetAll(http.ResponseWriter, *http.Request)
	GetAvailableBooks(http.ResponseWriter, *http.Request)
}

type BookHandler struct {
	Service *service.Manager
}

func (h *BookHandler) GetByAuthor(writer http.ResponseWriter, request *http.Request) {
	author := strings.TrimPrefix(request.URL.Path, "/api/books/getByAuthor/")
	res, _ := h.Service.Book.GetByAuthor(author)
	_ = json.NewEncoder(writer).Encode(res)
}

func (h *BookHandler) GetByName(writer http.ResponseWriter, request *http.Request) {
	name := strings.TrimPrefix(request.URL.Path, "/api/books/getByName/")
	res, _ := h.Service.Book.GetByName(name)
	_ = json.NewEncoder(writer).Encode(res)
}

func (h *BookHandler) Create(writer http.ResponseWriter, request *http.Request) {
	var createBook model.Book
	_ = json.NewDecoder(request.Body).Decode(&createBook)
	_ = h.Service.Book.Create(createBook)
	_, _ = writer.Write([]byte("Book created"))
}

func (h *BookHandler) Update(writer http.ResponseWriter, request *http.Request) {
	var updateBook model.Book
	_ = json.NewDecoder(request.Body).Decode(&updateBook)
	_ = h.Service.Book.Update(updateBook)
	_, _ = writer.Write([]byte("Book Updated"))
}

func (h *BookHandler) Delete(writer http.ResponseWriter, request *http.Request) {
	name := strings.TrimPrefix(request.URL.Path, "/api/books/")
	_ = h.Service.Book.Delete(name)
	_, _ = writer.Write([]byte("Book deleted"))
}

func (h *BookHandler) GetAll(writer http.ResponseWriter, request *http.Request) {
	res, _ := h.Service.Book.GetAll()
	_ = json.NewEncoder(writer).Encode(res)
}
func (h *BookHandler) GetAvailableBooks(writer http.ResponseWriter, request *http.Request) {
	res, _ := h.Service.Book.GetAvailableBooks()
	_ = json.NewEncoder(writer).Encode(res)
}

func NewBookHandler(s *service.Manager) *BookHandler {
	return &BookHandler{
		Service: s,
	}
}
