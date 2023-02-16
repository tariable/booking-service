package service

import (
	"bookingService/internal/adapter/repository"
	"bookingService/pkg/logger"
	"bookingService/pkg/metrics"
)

type OrderService struct {
	logger.Logger
	repository.OrderRepository
	metrics.Metrics
}

func NewOrderService(l logger.Logger, r repository.OrderRepository, m metrics.Metrics) *OrderService {
	return &OrderService{l, r, m}
}
