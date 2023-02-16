package entity

import (
	"bookingService/internal/dto"
	"time"
)

type Order struct {
	Room      string    `json:"room"`
	UserEmail string    `json:"email"`
	From      time.Time `json:"from"`
	To        time.Time `json:"to"`
}

func NewOrder(dto *dto.MakeOrder) *Order {
	from, _ := time.Parse("2006-01-02", dto.From)
	to, _ := time.Parse("2006-01-02", dto.To)

	order := &Order{
		Room:      dto.Room,
		UserEmail: dto.Email,
		From:      from,
		To:        to,
	}

	return order
}
