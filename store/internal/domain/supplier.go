package domain

import (
	"github.com/google/uuid"
	"time"
)

type Supplier struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	AddressID   *uuid.UUID `json:"address_id"`
	PhoneNumber string     `json:"phone_number"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
