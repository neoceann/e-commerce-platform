package repository

import (
	"context"
	"fmt"
	"store/internal/domain"
	"store/internal/dto"
	"store/internal/repository/db"
	"github.com/google/uuid"
)

type SupplierRepositoryImpl struct {
	queries *db.Queries
}

func NewSupplierRepository(q *db.Queries) SupplierRepository {
	return &SupplierRepositoryImpl{queries: q}
}

func (r *SupplierRepositoryImpl) CreateSupplier(ctx context.Context, request *dto.CreateSupplierRequest) (*domain.Supplier, error) {

	supplier, err := r.queries.CreateSupplier(ctx, db.CreateSupplierParams{
		Name: request.Name,
		PhoneNumber: request.PhoneNumber,
		AddressID: request.AddressID})

	if err != nil {
		return nil, fmt.Errorf("failed to create new supplier: %w", err)
	}

	return dto.SupplierFromDbToDomain(supplier), nil
}


func (r *SupplierRepositoryImpl) DeleteSupplier(ctx context.Context, supplierID uuid.UUID) error {
	return r.queries.DeleteSupplier(ctx, supplierID)
}

func (r *SupplierRepositoryImpl) GetAllSuppliers(ctx context.Context) ([]*domain.Supplier, error) {
	suppliers, err := r.queries.GetAllSuppliers(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to get suppliers: %w", err)
	}

	var domainSuppliers []*domain.Supplier

	for _, supplier := range suppliers {
		domainSuppliers = append(domainSuppliers, dto.SupplierFromDbToDomain(supplier))
	}

	return domainSuppliers, nil
}

func (r *SupplierRepositoryImpl) UpdateSupplierAddr(ctx context.Context, supplierID uuid.UUID, request *dto.UpdateAddressParamsRequest) (*domain.Supplier, error) {

	newAddr, err := r.queries.CreateAddress(ctx, db.CreateAddressParams{Country: request.Country,
													City: request.City,
													Street: request.Street})

	if err != nil {
		return nil, fmt.Errorf("failed to create address: %w", err)
	}

	supplier, err := r.queries.UpdateSupplierAddress(ctx, db.UpdateSupplierAddressParams{ID: supplierID, AddressID: newAddr.ID})


	if err != nil {
		return nil, fmt.Errorf("failed to update supplier address: %w", err)
	}

	return dto.SupplierFromDbToDomain(supplier), nil
}

func (r *SupplierRepositoryImpl) GetSupplierByID(ctx context.Context, supplierID uuid.UUID) (*domain.Supplier, error) {
	
	supplier, err := r.queries.GetSupplierByID(ctx, supplierID)

	if err != nil {
		return nil, fmt.Errorf("supplier not found")
	}

	return dto.SupplierFromDbToDomain(supplier), nil

}