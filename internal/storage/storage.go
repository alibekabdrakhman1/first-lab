package storage

import (
	"github.com/alibekabdrakhman1/first-lab/internal/model"
	"github.com/alibekabdrakhman1/first-lab/internal/storage/postgre"
)

type IUserRepository interface {
	Get(login string) (model.User, error)
	Create(user model.User) error
	Update(user model.User) error
	Delete(login string) error
	GetAll() ([]model.User, error)
}

type Storage struct {
	User IUserRepository
}

func New() (*Storage, error) {
	uRepo := postgre.NewUserRepository()
	var storage Storage
	storage.User = uRepo
	return &storage, nil
}
