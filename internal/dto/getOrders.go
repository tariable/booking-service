package dto

import (
	"fmt"
	"net/http"
)

type GetOrders struct {
	Email string
}

func NewGetOrders(r *http.Request) *GetOrders {
	dto := &GetOrders{
		Email: r.URL.Query().Get("email"),
	}

	return dto
}

func (dto *GetOrders) Validate() error {
	if dto.Email == "" {
		return fmt.Errorf("empty email")
	}

	return nil
}
