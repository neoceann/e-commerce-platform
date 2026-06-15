package service

import (
	"context"
	"store/internal/domain"
	"store/internal/dto"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, request *dto.CreateCategoryRequest) (*domain.Category, error)

	GetAllCategories(ctx context.Context) ([]*domain.Category, error)
}
