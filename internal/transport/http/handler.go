package http

import "github.com/alibekabdrakhman1/first-lab/internal/service"

type Handler struct {
	User IUserHandler
}

func NewHandler(service *service.Manager) *Handler {
	return &Handler{
		User: NewUserHandler(service),
	}
}
