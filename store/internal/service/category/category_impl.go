package service

import (
	"context"
	"fmt"
	"store/internal/domain"
	"store/internal/dto"
	"store/internal/repository/category"
	"strings"
)

type CategoryServiceImpl struct {
	clientRepo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{clientRepo: repo}
}

func (c *CategoryServiceImpl) CreateCategory(ctx context.Context, request *dto.CreateCategoryRequest) (*domain.Category, error) {
	if strings.TrimSpace(request.Name) == "" {
		return nil, fmt.Errorf("%w: category name is required", ErrInvalidCategoryData)
	}

	category, err := c.clientRepo.CreateCategory(ctx, request)
	if err != nil {
		return nil, err
	}

	return category, nil
}


func (c *CategoryServiceImpl) GetAllCategories(ctx context.Context) ([]*domain.Category, error) {
	categories, err := c.clientRepo.GetAllCategories(ctx)

	if err != nil {
		return nil, err
	}

	if categories == nil {
		return []*domain.Category{}, nil
	}

	return categories, nil
}