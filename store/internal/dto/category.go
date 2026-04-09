package dto

import (
	"time"
	"github.com/google/uuid"
	"store/internal/domain"
	"store/internal/repository/db"
)

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required"`
}

type CategoryResponce struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CategoryFromDbToDomain(c db.Category) *domain.Category {
	return &domain.Category{
		ID: c.ID,
		Name: c.Name,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}