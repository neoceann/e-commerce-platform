package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"store/internal/domain"
	"store/internal/dto"
	"store/internal/repository/product"
	"strings"
)

type ProductServiceImpl struct {
	productRepo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &ProductServiceImpl{productRepo: repo}
}

func (p *ProductServiceImpl) CreateProduct(ctx context.Context, request *dto.CreateProductRequest) (*domain.Product, error) {
	if strings.TrimSpace(request.Name) == "" {
		return nil, fmt.Errorf("%w: product name is required", ErrInvalidProductData)
	}

	if request.CategoryID == uuid.Nil {
		return nil, fmt.Errorf("%w: invalid category UUID", ErrInvalidID)
	}

	if request.SupplierID == uuid.Nil {
		return nil, fmt.Errorf("%w: invalid supplier UUID", ErrInvalidID)
	}

	if request.Price <= float64(0) {
		return nil, fmt.Errorf("%w: invalid price value (must be positive)", ErrBadPrice)
	}

	if request.AvailableStock < 0 {
		return nil, fmt.Errorf("%w: invalid stock value (must be non-negative)", ErrBadStock)
	}

	product, err := p.productRepo.CreateProduct(ctx, request)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductServiceImpl) IncreaseProductStock(ctx context.Context, id uuid.UUID, request *dto.IncreaseProductStockRequest) (*domain.Product, error) {

	if id == uuid.Nil {
		return nil, fmt.Errorf("%w: invalid product UUID", ErrInvalidID)
	}

	if request.Increasevalue <= 0 {
		return nil, fmt.Errorf("%w: invalid stock value (must be non-negative)", ErrIncreaseFailed)
	}

	_, err := p.productRepo.GetProductByID(ctx, id)

	if err != nil {
		return nil, ErrProductNotFound
	}

	product, err := p.productRepo.IncreaseProductStock(ctx, id, request)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductServiceImpl) DecreaseProductStock(ctx context.Context, id uuid.UUID, request *dto.DecreaseProductStockRequest) (*domain.Product, error) {
	if id == uuid.Nil {
		return nil, fmt.Errorf("%w: invalid product UUID", ErrInvalidID)
	}

	if request.Decreasevalue <= 0 {
		return nil, fmt.Errorf("%w: invalid stock value (must be non-negative)", ErrDecreaseFailed)
	}

	temp, err := p.productRepo.GetProductByID(ctx, id)

	if err != nil {
		return nil, ErrProductNotFound
	}

	if request.Decreasevalue > temp.AvailableStock {
		return nil, fmt.Errorf("%w: value can not be greater than available (%d)", ErrDecreaseFailed, temp.AvailableStock)
	}

	product, err := p.productRepo.DecreaseProductStock(ctx, id, request)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductServiceImpl) DeleteProductByID(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		return ErrInvalidID
	}

	_, err := p.productRepo.GetProductByID(ctx, id)

	if err != nil {
		return ErrProductNotFound
	}

	return p.productRepo.DeleteProductByID(ctx, id)
}

func (p *ProductServiceImpl) GetAvailableProducts(ctx context.Context) ([]*domain.Product, error) {
	products, err := p.productRepo.GetAvailableProducts(ctx)

	if err != nil {
		return nil, err
	}

	if products == nil {
		return []*domain.Product{}, nil
	}

	return products, nil
}

func (p *ProductServiceImpl) GetProductByID(ctx context.Context, id uuid.UUID) (*domain.Product, error) {
	if id == uuid.Nil {
		return nil, ErrInvalidID
	}

	product, err := p.productRepo.GetProductByID(ctx, id)

	if err != nil {
		return nil, ErrProductNotFound
	}

	return product, nil

}
