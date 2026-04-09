package repository

import (
	"context"
	"store/internal/domain"
	"store/internal/dto"
	"github.com/google/uuid"
)

type SupplierRepository interface {
	CreateSupplier(ctx context.Context, request *dto.CreateSupplierRequest) (*domain.Supplier, error)

	DeleteSupplier(ctx context.Context, supplierID uuid.UUID) error

	GetAllSuppliers(ctx context.Context) ([]*domain.Supplier, error)

	UpdateSupplierAddr(ctx context.Context, supplierID uuid.UUID, request *dto.UpdateAddressParamsRequest) (*domain.Supplier, error)

	GetSupplierByID(ctx context.Context, supplierID uuid.UUID) (*domain.Supplier, error)
}
