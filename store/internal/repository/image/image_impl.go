package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"store/internal/domain"
	"store/internal/dto"
	"store/internal/repository/db"
)

type ImageRepositoryImpl struct {
	queries *db.Queries
}

func NewImageRepository(q *db.Queries) ImageRepository {
	return &ImageRepositoryImpl{queries: q}
}

func (r *ImageRepositoryImpl) CreateImage(ctx context.Context, request *dto.CreateImageDataByte) (*domain.Image, error) {

	image, err := r.queries.CreateImage(ctx, db.CreateImageParams{ImageData: request.ImageData, ProductID: request.ProductID})

	if err != nil {
		return nil, fmt.Errorf("failed to create image: %w", err)
	}

	return dto.ImageFromDbToDomain(image), nil
}

func (r *ImageRepositoryImpl) UpdateImage(ctx context.Context, imageID uuid.UUID, request *dto.UpdateImageDataByte) (*domain.Image, error) {
	image, err := r.queries.UpdateImage(ctx, db.UpdateImageParams{ID: imageID, ImageData: request.ImageData})

	if err != nil {
		return nil, fmt.Errorf("failed to update image: %w", err)
	}

	return dto.ImageFromDbToDomain(image), nil
}

func (r *ImageRepositoryImpl) DeleteImageByID(ctx context.Context, imageID uuid.UUID) error {
	return r.queries.DeleteImageByID(ctx, imageID)
}

func (r *ImageRepositoryImpl) GetImagesByProductID(ctx context.Context, productID uuid.UUID) ([]*domain.Image, error) {
	images, err := r.queries.GetImageByProductID(ctx, productID)

	if err != nil {
		return nil, fmt.Errorf("failed to get image: %w", err)
	}

	var img []*domain.Image

	for _, image := range images {
		img = append(img, dto.ImageFromDbToDomain(image))
	}

	return img, nil
}

func (r *ImageRepositoryImpl) GetImageByImageId(ctx context.Context, imageID uuid.UUID) (*domain.Image, error) {
	image, err := r.queries.GetImageByImageId(ctx, imageID)

	if err != nil {
		return nil, fmt.Errorf("failed to get image: %w", err)
	}

	return dto.ImageFromDbToDomain(image), nil
}
