package repository

import (
	"bookingService/config"
	"bookingService/internal/entity"
	"bookingService/pkg/logger"
)

type OrderRepository interface {
	GetOrders() ([]entity.Order, error)
	AddOrder(order *entity.Order) error
}

func New(l logger.Logger, c config.Storage) (OrderRepository, error) {
	repo := &LocalStorage{
		logger: l,
		orders: make([]entity.Order, 0),
	}

	return repo, nil
}
