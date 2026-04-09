package repository

import (
	"context"
	"store/internal/domain"
	"store/internal/dto"
	"github.com/google/uuid"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, request *dto.CreateProductRequest) (*domain.Product, error)

	IncreaseProductStock(ctx context.Context, id uuid.UUID, request *dto.IncreaseProductStockRequest) (*domain.Product, error)

	DecreaseProductStock(ctx context.Context, id uuid.UUID, request *dto.DecreaseProductStockRequest) (*domain.Product, error)

	DeleteProductByID(ctx context.Context, id uuid.UUID) error

	GetAvailableProducts(ctx context.Context) ([]*domain.Product, error)

	GetProductByID(ctx context.Context, id uuid.UUID) (*domain.Product, error)
}
