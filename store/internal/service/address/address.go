package service

import (
	"context"
	"store/internal/domain"
	"github.com/google/uuid"
)

type AddressService interface {
	GetAddressByID(ctx context.Context, id uuid.UUID) (*domain.Address, error)
}