package http

import (
	"encoding/json"
	"github.com/alibekabdrakhman1/first-lab/internal/model"
	"github.com/alibekabdrakhman1/first-lab/internal/service"
	"net/http"
	"strings"
)

type IUserHandler interface {
	Get(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	GetAll(http.ResponseWriter, *http.Request)
}

type UserHandler struct {
	Service *service.Manager
}

func NewUserHandler(s *service.Manager) *UserHandler {
	return &UserHandler{
		Service: s,
	}
}

func (h *UserHandler) Get(writer http.ResponseWriter, request *http.Request) {
	login := strings.TrimPrefix(request.URL.Path, "/api/users/")
	res, _ := h.Service.User.Get(login)
	json.NewEncoder(writer).Encode(res)
}

func (h *UserHandler) Create(writer http.ResponseWriter, request *http.Request) {
	var createUser model.User
	json.NewDecoder(request.Body).Decode(&createUser)
	h.Service.User.Create(createUser)
	writer.Write([]byte("User created"))
}

func (h *UserHandler) Update(writer http.ResponseWriter, request *http.Request) {

}

func (h *UserHandler) Delete(writer http.ResponseWriter, request *http.Request) {
	login := strings.TrimPrefix(request.URL.Path, "/api/users/")
	h.Service.User.Delete(login)
	writer.Write([]byte("User deleted"))
}

func (h *UserHandler) GetAll(writer http.ResponseWriter, request *http.Request) {
	res, _ := h.Service.User.GetAll()
	json.NewEncoder(writer).Encode(res)
}
