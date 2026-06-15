package repository

import (
	"context"
	"github.com/google/uuid"
	"store/internal/domain"
)

type AddressRepository interface {
	GetAddressByID(ctx context.Context, id uuid.UUID) (*domain.Address, error)
}
