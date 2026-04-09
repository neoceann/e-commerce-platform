package repository

import (
	"context"
	"store/internal/domain"
	"github.com/google/uuid"
)

type AddressRepository interface {
	GetAddressByID(ctx context.Context, id uuid.UUID) (*domain.Address, error)
}