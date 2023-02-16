package repository

import (
	"bookingService/internal/entity"
	"bookingService/pkg/logger"
	"sync"
)

type LocalStorage struct {
	logger logger.Logger
	mutex  sync.RWMutex
	orders []entity.Order
}

func (s *LocalStorage) GetOrders() ([]entity.Order, error) {
	s.mutex.RLock()
	o := s.orders
	s.mutex.RUnlock()

	return o, nil
}

func (s *LocalStorage) AddOrder(order *entity.Order) error {
	s.mutex.Lock()
	s.orders = append(s.orders, *order)
	s.mutex.Unlock()

	return nil
}
