package dto

import (
	"store/internal/domain"
	"store/internal/repository/db"
)

type UpdateAddressParamsRequest struct {
	Country string    `json:"country" validate:"required"`
	City    string    `json:"city" validate:"required"`
	Street  string    `json:"street" validate:"required"`
}

func AddressFromDbToDomain(a db.Address) *domain.Address {
	return &domain.Address{
		ID: a.ID,
		Country: a.Country,
		City: a.City,
		Street: a.Street,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}