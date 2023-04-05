package transport

import (
	"net/http"

	"github.com/alibekabdrakhman1/first-lab/internal/service"
)
// это должно быть в htpp/user.go
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

func (h *UserHandler) Get(writer http.ResponseWriter, request *http.Request) {
	panic("implement me")
}

func (h *UserHandler) Create(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h *UserHandler) Update(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h *UserHandler) Delete(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h *UserHandler) GetAll(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func NewUserHandler(s *service.Manager) *UserHandler {
	return &UserHandler{
		Service: s,
	}
}
