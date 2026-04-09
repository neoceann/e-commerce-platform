package service

import (
	"context"
	"store/internal/dto"
	"store/internal/domain"
	"github.com/google/uuid"
)

type SupplierService interface {
	CreateSupplier(ctx context.Context, request *dto.CreateSupplierRequest) (*domain.Supplier, error)

	DeleteSupplier(ctx context.Context, supplierID uuid.UUID) error

	GetAllSuppliers(ctx context.Context) ([]*domain.Supplier, error)

	GetSupplierByID(ctx context.Context, supplierID uuid.UUID) (*domain.Supplier, error)

	UpdateSupplierAddr(ctx context.Context, supplierID uuid.UUID, request *dto.UpdateAddressParamsRequest) (*domain.Supplier, error)
}