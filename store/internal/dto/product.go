package dto

import (
	"github.com/google/uuid"
	"store/internal/domain"
	"store/internal/repository/db"
	"time"
)

type CreateProductRequest struct {
	Name           string    `json:"name" validate:"required"`
	CategoryID     uuid.UUID `json:"category_id" validate:"required"`
	Price          float64   `json:"price" validate:"required"`
	AvailableStock int16     `json:"available_stock" validate:"required"`
	SupplierID     uuid.UUID `json:"supplier_id" validate:"required"`
}

type IncreaseProductStockRequest struct {
	Increasevalue int16 `json:"increasevalue"`
}

type DecreaseProductStockRequest struct {
	Decreasevalue int16 `json:"decreasevalue"`
}

type ProductResponce struct {
	ID                      uuid.UUID `json:"id"`
	Name                    string    `json:"name"`
	CategoryID              uuid.UUID `json:"category_id"`
	Price                   float64   `json:"price"`
	AvailableStock          int16     `json:"available_stock"`
	LastAvailableUpdateDate time.Time `json:"last_available_update_date"`
	SupplierID              uuid.UUID `json:"supplier_id"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
}

func ProductFromDbToDomain(p db.Product) *domain.Product {
	return &domain.Product{
		ID:                      p.ID,
		Name:                    p.Name,
		CategoryID:              p.CategoryID,
		Price:                   p.Price,
		AvailableStock:          p.AvailableStock,
		LastAvailableUpdateDate: p.LastAvailableUpdateDate,
		SupplierID:              p.SupplierID,
		CreatedAt:               p.CreatedAt,
		UpdatedAt:               p.UpdatedAt,
	}
}
