package service

import (
	"context"
	"fmt"
	"store/internal/domain"
	"store/internal/dto"
	"store/internal/repository/supplier"
	"strings"
	"github.com/google/uuid"
)

type SupplierServiceImpl struct {
	SupplierRepo repository.SupplierRepository
}

func NewSupplierService(repo repository.SupplierRepository) SupplierService {
	return &SupplierServiceImpl{SupplierRepo: repo}
}

func (s *SupplierServiceImpl) CreateSupplier(ctx context.Context, request *dto.CreateSupplierRequest) (*domain.Supplier, error) {
	if strings.TrimSpace(request.Name) == "" {
		return nil, fmt.Errorf("%w: supplier name is required", ErrInvalidSupplierData)
	}

	supplier, err := s.SupplierRepo.CreateSupplier(ctx, request)
	if err != nil {
		return nil, err
	}

	return supplier, nil
}

func (s *SupplierServiceImpl) GetAllSuppliers(ctx context.Context) ([]*domain.Supplier, error) {
	suppliers, err := s.SupplierRepo.GetAllSuppliers(ctx)

	if err != nil {
		return nil, err
	}

	if suppliers == nil {
		return []*domain.Supplier{}, nil
	}

	return suppliers, nil
}

func (s *SupplierServiceImpl) GetSupplierByID(ctx context.Context, supplierID uuid.UUID) (*domain.Supplier, error) {
	if supplierID == uuid.Nil {
		return nil, ErrInvalidID
	}

	supplier, err := s.SupplierRepo.GetSupplierByID(ctx, supplierID)

	if err != nil {
		return nil, ErrSupplierNotFound
	}

	return supplier, nil
}

func (s *SupplierServiceImpl) DeleteSupplier(ctx context.Context, supplierID uuid.UUID) error {
	if supplierID == uuid.Nil {
		return ErrInvalidID
	}

	_, err := s.SupplierRepo.GetSupplierByID(ctx, supplierID)

	if err != nil {
		return ErrSupplierNotFound
	}

	return s.SupplierRepo.DeleteSupplier(ctx, supplierID)
}

func (s *SupplierServiceImpl) UpdateSupplierAddr(ctx context.Context, supplierID uuid.UUID, request *dto.UpdateAddressParamsRequest) (*domain.Supplier, error) {
	if supplierID == uuid.Nil {
		return nil, ErrInvalidID
	}

	if strings.TrimSpace(request.Country) == "" ||
		strings.TrimSpace(request.City) == "" ||
		strings.TrimSpace(request.Street) == "" {
			return nil, ErrInvalidAddrData
	}

	_, err := s.SupplierRepo.GetSupplierByID(ctx, supplierID)

	if err != nil {
		return nil, ErrSupplierNotFound
	}

	return s.SupplierRepo.UpdateSupplierAddr(ctx, supplierID, request)
}