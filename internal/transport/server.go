package transport

import (
	"github.com/alibekabdrakhman1/first-lab/internal/service"
)

type Handler struct {
	User IUserHandler
}

func NewHandler() *Handler {
	return &Handler{
		User: NewUserHandler(service.NewManager()),
	}
}