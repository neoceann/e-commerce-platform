package service

import (
	"context"
	"github.com/google/uuid"
	"store/internal/domain"
	"store/internal/dto"
)

type ImageService interface {
	CreateImage(ctx context.Context, request *dto.CreateImageRequest) (*domain.Image, error)
	UpdateImage(ctx context.Context, imageID uuid.UUID, request *dto.UpdateImageRequest) (*domain.Image, error)
	DeleteImageByID(ctx context.Context, imageID uuid.UUID) error
	GetImagesByProductID(ctx context.Context, productID uuid.UUID) ([]*domain.Image, error)
	GetImageByImageId(ctx context.Context, imageID uuid.UUID) (*domain.Image, error)
}
