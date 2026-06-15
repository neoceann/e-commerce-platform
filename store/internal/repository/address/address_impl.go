package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"store/internal/domain"
	"store/internal/dto"
	"store/internal/repository/db"
)

type AddressRepositoryImpl struct {
	queries *db.Queries
}

func NewAddressRepository(q *db.Queries) AddressRepository {
	return &AddressRepositoryImpl{queries: q}
}

func (r *AddressRepositoryImpl) GetAddressByID(ctx context.Context, id uuid.UUID) (*domain.Address, error) {
	address, err := r.queries.GetAddressByID(ctx, id)

	if err != nil {
		return nil, fmt.Errorf("failed to get address: %w", err)
	}

	return dto.AddressFromDbToDomain(address), nil
}
