package service

import (
	"github.com/alibekabdrakhman1/first-lab/internal/model"
	"github.com/alibekabdrakhman1/first-lab/internal/storage"
	"github.com/google/uuid"
)

type IUserService interface {
	Get(login string) (model.User, error)
	Create(user model.User) error
	Update(user model.User) error
	Delete(login string) error
	GetAll() ([]model.User, error)
}

type UserService struct {
	repo *storage.Storage
}

func NewUserService(r *storage.Storage) *UserService {
	return &UserService{
		repo: r,
	}
}

func (s *UserService) Get(login string) (model.User, error) {
	return s.repo.User.Get(login)
}

func (s *UserService) Create(user model.User) error {
	user.ID = uuid.New().String()
	return s.repo.User.Create(user)
}

func (s *UserService) Update(user model.User) error {
	return s.repo.User.Update(user)
}
func (s *UserService) Delete(login string) error {
	return s.repo.User.Delete(login)
}

func (s *UserService) GetAll() ([]model.User, error) {
	return s.repo.User.GetAll()
}
