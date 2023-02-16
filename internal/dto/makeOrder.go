package dto

import (
	"fmt"
	"net/http"
	"time"
)

var AvailableRooms = map[string]struct{}{"econom": {}, "standart": {}, "lux": {}}

type MakeOrder struct {
	Room  string
	Email string
	From  string
	To    string
}

func NewMakeOrder(r *http.Request) *MakeOrder {
	dto := &MakeOrder{
		Room:  r.URL.Query().Get("room"),
		Email: r.URL.Query().Get("email"),
		From:  r.URL.Query().Get("from"),
		To:    r.URL.Query().Get("to"),
	}

	return dto
}

func (dto *MakeOrder) Validate() error {
	if dto.Email == "" {
		return fmt.Errorf("empty email")
	}

	if dto.Room == "" {
		return fmt.Errorf("empty room")
	}
	if _, exists := AvailableRooms[dto.Room]; !exists {
		return fmt.Errorf("invalid room")
	}

	if dto.From == "" {
		return fmt.Errorf("empty from")
	}
	to, err := time.Parse("2006-01-02", dto.From)
	if err != nil {
		return fmt.Errorf("invalid format of from")
	}

	if dto.To == "" {
		return fmt.Errorf("empty from")
	}
	from, err := time.Parse("2006-01-02", dto.To)
	if err != nil {
		return fmt.Errorf("invalid format of from")
	}

	if to.After(from) || from.Before(time.Now()) {
		return fmt.Errorf("invalid time period")
	}

	return nil
}
