package postgre

import (
	"errors"

	"github.com/alibekabdrakhman1/first-lab/internal/model"
)

type UserRepository struct {
	db []model.User
}

func (r *UserRepository) Get(login string) (model.User, error) {
	for _, user := range r.db {
		if user.Login == login {
			return user, nil
		}
	}
	return model.User{}, errors.New("user not found")
}

func (r *UserRepository) Create(user model.User) error {
	user.Id = len(r.db) + 1
	r.db = append(r.db, user)
	return nil
}

func (r *UserRepository) Update(user model.User) error {
	for i := 0; i < len(r.db); i++ {
		if r.db[i].Id == user.Id {
			r.db[i] = user
			return nil
		}
	}
	return errors.New("user not found")
}

func (r *UserRepository) Delete(login string) error {
	for i := 0; i < len(r.db); i++ {
		if r.db[i].Login == login {
			r.db = remove(r.db, i)
			return nil
		}
	}
	return errors.New("user not found")
}

func (r *UserRepository) GetAll() ([]model.User, error) {
	return r.db, nil
}

func remove(slice []model.User, index int) []model.User {
	return append(slice[:index], slice[index+1:]...)
}
func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: []model.User{},
	}
}
