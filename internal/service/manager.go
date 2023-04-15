package service

import "github.com/alibekabdrakhman1/first-lab/internal/storage"

type Manager struct {
	User  IUserService
	Book  IBookService
	Order IOrderService
}

func NewManager(storage *storage.Storage) (*Manager, error) {
	uSrv := NewUserService(storage)
	bSrv := NewBookService(storage)
	oSrv := NewOrderService(storage)

	return &Manager{
		User:  uSrv,
		Book:  bSrv,
		Order: oSrv,
	}, nil
}
