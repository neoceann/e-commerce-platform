package service

import (
	"context"
	"github.com/google/uuid"
	"store/internal/domain"
	"store/internal/repository/address"
)

type AddressServiceImpl struct {
	addressRepo repository.AddressRepository
}

func NewAddressService(r repository.AddressRepository) AddressService {
	return &AddressServiceImpl{addressRepo: r}
}

func (s *AddressServiceImpl) GetAddressByID(ctx context.Context, id uuid.UUID) (*domain.Address, error) {
	if id == uuid.Nil {
		return nil, ErrInvalidID
	}

	address, err := s.addressRepo.GetAddressByID(ctx, id)

	if address == nil {
		return nil, ErrAddrNotFound
	}

	if err != nil {
		return nil, err
	}

	return address, nil
}
