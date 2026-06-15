package dto

import (
	"github.com/google/uuid"
	"store/internal/domain"
	"store/internal/repository/db"
	"time"
)

type CreateClientRequest struct {
	ClientName    string     `json:"client_name" validate:"required"`
	ClientSurname string     `json:"client_surname" validate:"required"`
	Birthday      time.Time  `json:"birthday,omitempty"`
	Gender        string     `json:"gender,omitempty"`
	AddressID     *uuid.UUID `json:"address_id,omitempty"`
}

type GetClientsByNameRequest struct {
	Name    string `json:"client_name"`
	SurName string `json:"client_surname"`
}

type GetclientsWithPaginationRequest struct {
	Limit  *int32 `json:"limit"`
	Offset *int32 `json:"offset"`
}

type ClientResponse struct {
	ID               uuid.UUID  `json:"id"`
	ClientName       string     `json:"client_name"`
	ClientSurname    string     `json:"client_surname"`
	Birthday         time.Time  `json:"birthday,omitempty"`
	Gender           string     `json:"gender,omitempty"`
	RegistrationDate time.Time  `json:"registration_date"`
	AddressID        *uuid.UUID `json:"address_id,omitempty"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}

func ClientFromDbToDomain(c db.Client) *domain.Client {
	return &domain.Client{
		ID:               c.ID,
		ClientName:       c.ClientName,
		ClientSurname:    c.ClientSurname,
		Birthday:         c.Birthday,
		Gender:           c.Gender,
		AddressID:        &c.AddressID,
		RegistrationDate: c.RegistrationDate,
		UpdatedAt:        c.UpdatedAt,
	}
}
