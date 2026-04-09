package service

import (
	"context"
	"store/internal/dto"
	"store/internal/domain"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, request *dto.CreateCategoryRequest) (*domain.Category, error)

	GetAllCategories(ctx context.Context) ([]*domain.Category, error)
}