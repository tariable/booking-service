package service

import (
	"bookingService/internal/dto"
	"bookingService/internal/entity"
	"net/http"
)

func (s *OrderService) MakeOrder(w http.ResponseWriter, r *http.Request) {
	// create DTO and validate it
	orderDto := dto.NewMakeOrder(r)
	err := orderDto.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	// creation of entity
	order := entity.NewOrder(orderDto)

	// availability check
	orders, err := s.OrderRepository.GetOrders()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// check type of room and time intervals
	for _, existedOrder := range orders {
		if order.Room == existedOrder.Room &&
			(((order.From.Equal(existedOrder.From) || order.From.After(existedOrder.From)) && order.From.Before(existedOrder.To)) ||
				((order.To.Equal(existedOrder.To) || order.From.Before(existedOrder.From)) && order.To.After(existedOrder.From))) {
			http.Error(w, "dates are not available", http.StatusConflict)
			return
		}
	}

	// store valid order in storage
	err = s.OrderRepository.AddOrder(order)
	if err != nil {
		s.Logger.Errorf("store order: %s", err.Error())
		return
	}

	// response, metrics, logs
	s.Logger.Info("Method makeOrder was successfully done")
	s.Metrics.IncreaseCounter("make-order")
	w.WriteHeader(http.StatusCreated)
}
