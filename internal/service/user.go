package service

import (
	"github.com/alibekabdrakhman1/first-lab/internal/model"
	"github.com/alibekabdrakhman1/first-lab/internal/repository"
)

type IUserService interface {
	Get(id int) (model.User, error)
	Create(user model.User) error
	Update(user model.User) error
	Delete(id int) error
	GetAll() ([]model.User, error)
}
// нету конструктора
type UserService struct {
	repo repository.Repository
}