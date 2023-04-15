package postgre

import (
	"github.com/alibekabdrakhman1/first-lab/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (r *UserRepository) Get(login string) (model.User, error) {
	var res model.User
	err := r.db.Where("login = ?", login).First(&res).Error
	if err != nil {
		return model.User{}, err
	}
	return res, nil
}

func (r *UserRepository) Create(user model.User) error {
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Update(user model.User) error {
	return r.db.Where("login = ", user.Login).Updates(&user).Error
}

func (r *UserRepository) Delete(login string) error {
	return r.db.Delete(&model.User{}, login).Error

}

func (r *UserRepository) GetAll() ([]model.User, error) {
	var resp []model.User
	err := r.db.Find(&resp)
	return resp, err.Error
}

func NewUserRepository(DB *gorm.DB) *UserRepository {
	return &UserRepository{
		db: DB,
	}
}
