package repository

import (
	"context"
	"fmt"
	"store/internal/domain"
	"store/internal/dto"
	"store/internal/repository/db"
	"github.com/google/uuid"
)

type ProductRepositoryImpl struct {
	queries *db.Queries
}

func NewProductRepository(q *db.Queries) ProductRepository {
	return &ProductRepositoryImpl{queries: q}
}

func (p *ProductRepositoryImpl) CreateProduct(ctx context.Context, request *dto.CreateProductRequest) (*domain.Product, error) {

	product, err := p.queries.CreateProduct(ctx, db.CreateProductParams{
		Name: request.Name,
		CategoryID: request.CategoryID,
		Price: request.Price,
		AvailableStock: request.AvailableStock,
		SupplierID: request.SupplierID,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create new product: %w", err)
	}

	return dto.ProductFromDbToDomain(product), nil
}

func (p *ProductRepositoryImpl) IncreaseProductStock(ctx context.Context, id uuid.UUID, request *dto.IncreaseProductStockRequest) (*domain.Product, error){
	product, err := p.queries.IncreaseProductStock(ctx, db.IncreaseProductStockParams{ID: id, Increasevalue: request.Increasevalue})

	if err != nil {
		return nil, fmt.Errorf("failed to increase product stock: %w", err)
	}

	return dto.ProductFromDbToDomain(product), nil

}

func (p *ProductRepositoryImpl) DecreaseProductStock(ctx context.Context, id uuid.UUID, request *dto.DecreaseProductStockRequest) (*domain.Product, error){
	product, err := p.queries.DecreaseProductStock(ctx, db.DecreaseProductStockParams{ID: id, Decreasevalue: request.Decreasevalue})

	if err != nil {
		return nil, fmt.Errorf("failed to decrease product stock: %w", err)
	}

	return dto.ProductFromDbToDomain(product), nil
}

func (p *ProductRepositoryImpl) DeleteProductByID(ctx context.Context, id uuid.UUID) error{
	return p.queries.DeleteProductByID(ctx, id)
}

func (p *ProductRepositoryImpl) GetAvailableProducts(ctx context.Context) ([]*domain.Product, error){
	productList, err := p.queries.GetAvailableProducts(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to get products: %w", err)
	}

	var products []*domain.Product

	for _, product := range productList {
		products = append(products, dto.ProductFromDbToDomain(product))
	}

	return products, nil
}

func (p *ProductRepositoryImpl) GetProductByID(ctx context.Context, id uuid.UUID) (*domain.Product, error){
	product, err := p.queries.GetProductByID(ctx, id)

	if err != nil {
		return nil, fmt.Errorf("failed to get product: %w", err)
	}

	return dto.ProductFromDbToDomain(product), nil
}