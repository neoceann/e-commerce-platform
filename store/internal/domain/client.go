package domain

import (
	"github.com/google/uuid"
	"time"
)

type Client struct {
	ID               uuid.UUID  `json:"id"`
	ClientName       string     `json:"client_name"`
	ClientSurname    string     `json:"client_surname"`
	Birthday         time.Time  `json:"birthday"`
	Gender           string     `json:"gender"`
	RegistrationDate time.Time  `json:"registration_date"`
	AddressID        *uuid.UUID `json:"address_id"`
	UpdatedAt        time.Time  `json:"updated_at"`
}
