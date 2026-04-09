package domain

import (
	"time"
	"github.com/google/uuid"
)

type Product struct {
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