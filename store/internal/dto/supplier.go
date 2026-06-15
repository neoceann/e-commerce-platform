package dto

import (
	"store/internal/domain"
	"store/internal/repository/db"
	"time"

	"github.com/google/uuid"
)

type CreateSupplierRequest struct {
	Name        string     `json:"name" validate:"required"`
	AddressID   *uuid.UUID `json:"address_id,omitempty"`
	PhoneNumber string     `json:"phone_number,omitempty"`
}

type SupplierResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	AddressID   uuid.UUID `json:"address_id,omitempty"`
	PhoneNumber string    `json:"phone_number,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func SupplierFromDbToDomain(s db.Supplier) *domain.Supplier {
	return &domain.Supplier{
		ID:          s.ID,
		Name:        s.Name,
		AddressID:   &s.AddressID,
		PhoneNumber: s.PhoneNumber,
		CreatedAt:   s.CreatedAt,
		UpdatedAt:   s.UpdatedAt,
	}
}
