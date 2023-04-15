package storage

import (
	"github.com/alibekabdrakhman1/first-lab/configs"
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
type IBookRepository interface {
	GetByAuthor(author string) ([]model.Book, error)
	GetByName(name string) (model.Book, error)
	Create(book model.Book) error
	Update(book model.Book) error
	Delete(name string) error
	GetAll() ([]model.Book, error)
	GetAvailableBooks() ([]model.Book, error)
}
type IOrderRepository interface {
	GetAll() ([]model.Order, error)
	GetNotReturned() ([]model.Order, error)
	GetCountOrders() ([]model.Order, error)
	Create(order model.Order) error
	Delete(orderId string) error
	ReturnBook(orderID string) error
}
type Storage struct {
	User  IUserRepository
	Book  IBookRepository
	Order IOrderRepository
}

func New(cfg *configs.Config) (*Storage, error) {
	db, err := postgre.Dial(cfg)
	if err != nil {
		return nil, err
	}
	uRepo := postgre.NewUserRepository(db)
	bRepo := postgre.NewBookRepository(db)
	oRepo := postgre.NewOrderRepository(db)
	var storage Storage
	storage.User = uRepo
	storage.Book = bRepo
	storage.Order = oRepo
	return &storage, nil
}
