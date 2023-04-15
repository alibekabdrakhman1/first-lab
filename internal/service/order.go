package service

import (
	"github.com/alibekabdrakhman1/first-lab/internal/model"
	"github.com/alibekabdrakhman1/first-lab/internal/storage"
	"github.com/google/uuid"
	"time"
)

type IOrderService interface {
	GetAll() ([]model.Order, error)
	GetNotReturned() ([]model.Order, error)
	GetCountOrders() ([]model.Order, error)
	Create(order model.Order) error
	Delete(orderId string) error
	ReturnBook(orderId string) error
}

type OrderService struct {
	repo *storage.Storage
}

func (s *OrderService) GetAll() ([]model.Order, error) {
	return s.repo.Order.GetAll()
}

func (s *OrderService) GetNotReturned() ([]model.Order, error) {
	return s.repo.Order.GetNotReturned()
}

func (s *OrderService) GetCountOrders() ([]model.Order, error) {
	return s.repo.Order.GetCountOrders()
}

func (s *OrderService) Create(order model.Order) error {
	order.OrderedDate = time.Now()
	order.ID = uuid.New().String()
	return s.repo.Order.Create(order)
}

func (s *OrderService) Delete(orderId string) error {
	return s.repo.Order.Delete(orderId)
}
func (s *OrderService) ReturnBook(orderId string) error {
	return s.repo.Order.ReturnBook(orderId)
}

func NewOrderService(r *storage.Storage) *OrderService {
	return &OrderService{
		repo: r,
	}
}
