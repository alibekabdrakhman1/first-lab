package postgre

import (
	"github.com/alibekabdrakhman1/first-lab/internal/model"
	"gorm.io/gorm"
	"time"
)

type OrderRepository struct {
	db *gorm.DB
}

func (r *OrderRepository) Delete(orderId string) error {
	return r.db.Delete(&model.Order{}, orderId).Error
}

func (r *OrderRepository) ReturnBook(orderID string) error {
	return r.db.Model(model.Order{ID: orderID}).Update("returned_date", time.Now()).Error
}

func (r *OrderRepository) GetAll() ([]model.Order, error) {
	var all []model.Order
	if err := r.db.Preload("Book").Preload("User").Find(&all).Error; err != nil {
		return nil, err
	}
	return all, nil
}

func (r *OrderRepository) GetNotReturned() ([]model.Order, error) {
	var notReturned []model.Order
	if err := r.db.Where("returned is null").Find(&notReturned).Error; err != nil {
		return nil, err
	}
	return notReturned, nil
}

func (r *OrderRepository) GetCountOrders() ([]model.Order, error) {
	var orders []model.Order
	r.db.Preload("Book").Preload("User").Where("returned_date >= NOW() - INTERVAL '1 MONTH'").Find(&orders)
	return orders, nil
}

func (r *OrderRepository) Create(order model.Order) error {
	if err := r.db.Create(&order).Error; err != nil {
		return err
	}
	return nil
}

func NewOrderRepository(DB *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: DB,
	}
}
