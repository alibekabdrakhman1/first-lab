package transport

import (
	"github.com/alibekabdrakhman1/first-lab/internal/service"
)

type Handler struct {
	User IUserHandler
}
// файл сервер для инициализация сервера, а у тебя инициализация хендлера 
// эти файлы должны лежать внутри папки http 
func NewHandler() *Handler {
	return &Handler{
		User: NewUserHandler(service.NewManager()),
	}
}