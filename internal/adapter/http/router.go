package http

import (
	"bookingService/internal/service"
	"net/http"
)

func NewRouter(s *service.OrderService) *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/order", s.MakeOrder)
	router.HandleFunc("/orders", s.GetOrders)

	return router
}
