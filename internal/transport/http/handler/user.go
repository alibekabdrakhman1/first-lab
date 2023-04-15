package handler

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
	_ = h.Service.User.Create(createUser)
	_, _ = writer.Write([]byte("User created"))
}

func (h *UserHandler) Update(writer http.ResponseWriter, request *http.Request) {
	var updateUser model.User
	json.NewDecoder(request.Body).Decode(&updateUser)
	_ = h.Service.User.Update(updateUser)
	_, _ = writer.Write([]byte("User updated"))
}

func (h *UserHandler) Delete(writer http.ResponseWriter, request *http.Request) {
	login := strings.TrimPrefix(request.URL.Path, "/api/users/")
	_ = h.Service.User.Delete(login)
	_, _ = writer.Write([]byte("User deleted"))
}

func (h *UserHandler) GetAll(writer http.ResponseWriter, request *http.Request) {
	res, _ := h.Service.User.GetAll()
	_ = json.NewEncoder(writer).Encode(res)
}
