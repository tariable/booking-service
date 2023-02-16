package service

import (
	"bookingService/internal/dto"
	"bookingService/internal/entity"
	"encoding/json"
	"net/http"
)

func (s *OrderService) GetOrders(w http.ResponseWriter, r *http.Request) {
	// create DTO and validate it
	getOrdersDto := dto.NewGetOrders(r)
	err := getOrdersDto.Validate()
	if err != nil {
		s.Logger.Errorf("Validation error: %s", err)
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	result := make([]entity.Order, 0)
	orders, err := s.OrderRepository.GetOrders()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, o := range orders {
		if o.UserEmail == getOrdersDto.Email {
			result = append(result, o)
		}
	}

	b, err := json.Marshal(result)
	if err != nil {
		s.Logger.Errorf("error - get order by email: %s", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	_, err = w.Write(b)
	if err != nil {
		s.Logger.Errorf("error - get order by email: %s", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	s.Logger.Info("Method getOrders was successfully done")
	s.Metrics.IncreaseCounter("get-orders")
}
