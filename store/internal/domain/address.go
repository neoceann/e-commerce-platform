package domain

import (
	"github.com/google/uuid"
	"time"
)

type Address struct {
	ID        uuid.UUID `json:"id"`
	Country   string    `json:"country"`
	City      string    `json:"city"`
	Street    string    `json:"street"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
