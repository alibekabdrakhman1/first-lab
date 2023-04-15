package handler

import (
	"github.com/alibekabdrakhman1/first-lab/internal/service"
)

type Handler struct {
	User  IUserHandler
	Book  IBookHandler
	Order IOrderHandler
}

func NewHandler(service *service.Manager) *Handler {
	return &Handler{
		User:  NewUserHandler(service),
		Book:  NewBookHandler(service),
		Order: NewOrderHandler(service),
	}
}
