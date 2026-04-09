package repository

import (
	"context"
	"fmt"
	"store/internal/domain"
	"store/internal/dto"
	"store/internal/repository/db"
)

type CategoryRepositoryImpl struct {
	queries *db.Queries
}

func NewCategoryRepositry(q *db.Queries) CategoryRepository {
	return &CategoryRepositoryImpl{queries: q}
}

func (c *CategoryRepositoryImpl) CreateCategory(ctx context.Context, request *dto.CreateCategoryRequest) (*domain.Category, error) {
	category, err := c.queries.CreateCategory(ctx, request.Name)

	if err != nil {
		return nil, fmt.Errorf("failed to create new category: %w", err)
	}

	return dto.CategoryFromDbToDomain(category), nil
}

func (c *CategoryRepositoryImpl) GetAllCategories(ctx context.Context) ([]*domain.Category, error) {
	categoryList, err := c.queries.GetAllCategories(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to get categories: %w", err)
	}

	var categories []*domain.Category

	for _, category := range categoryList {
		categories = append(categories, dto.CategoryFromDbToDomain(category))
	}

	return categories, nil
}