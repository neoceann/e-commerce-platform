package service

import (
	"context"
	"store/internal/dto"
	"store/internal/domain"
	"github.com/google/uuid"
)

type ProductService interface {
	CreateProduct(ctx context.Context, request *dto.CreateProductRequest) (*domain.Product, error)

	IncreaseProductStock(ctx context.Context, id uuid.UUID, request *dto.IncreaseProductStockRequest) (*domain.Product, error)

	DecreaseProductStock(ctx context.Context, id uuid.UUID, request *dto.DecreaseProductStockRequest) (*domain.Product, error)

	DeleteProductByID(ctx context.Context, id uuid.UUID) error

	GetAvailableProducts(ctx context.Context) ([]*domain.Product, error)

	GetProductByID(ctx context.Context, id uuid.UUID) (*domain.Product, error)
}