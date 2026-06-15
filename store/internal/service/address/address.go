package service

import (
	"context"
	"github.com/google/uuid"
	"store/internal/domain"
)

type AddressService interface {
	GetAddressByID(ctx context.Context, id uuid.UUID) (*domain.Address, error)
}
