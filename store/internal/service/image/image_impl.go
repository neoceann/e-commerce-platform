package service

import (
	"context"
	"fmt"
	"log"
	"store/internal/domain"
	"store/internal/dto"
	"store/internal/repository/image"

	"github.com/google/uuid"
)

type ImageServiceImpl struct {
	imageRepo repository.ImageRepository
}

func NewImageService(r repository.ImageRepository) ImageService {
	return &ImageServiceImpl{imageRepo: r}
}

func (r *ImageServiceImpl) CreateImage(ctx context.Context, request *dto.CreateImageRequest) (*domain.Image, error) {

	if len(request.ImageData) == 0 {
		return nil, ErrInvalidImageData
	}

	imageData, err := dto.StringToBytes(request.ImageData)
	log.Printf("Received image data size: %d bytes", len(imageData))

	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidImageData, err)
	}

	image, err := r.imageRepo.CreateImage(ctx, &dto.CreateImageDataByte{ImageData: imageData, ProductID: request.ProductID})

	if err != nil {
		return nil, err
	}

	return image, nil
}

func (r *ImageServiceImpl) UpdateImage(ctx context.Context, imageID uuid.UUID, request *dto.UpdateImageRequest) (*domain.Image, error) {
	if imageID == uuid.Nil {
		return nil, ErrInvalidID
	}

	if len(request.ImageData) == 0 {
		return nil, ErrInvalidImageData
	}

	imageData, err := dto.StringToBytes(request.ImageData)

	image, err := r.imageRepo.UpdateImage(ctx, imageID, &dto.UpdateImageDataByte{ImageData: imageData})

	if err != nil {
		return nil, ErrImageNotFound
	}

	return image, nil
}

func (r *ImageServiceImpl) DeleteImageByID(ctx context.Context, imageID uuid.UUID) error {
	if imageID == uuid.Nil {
		return ErrInvalidID
	}

	_, err := r.imageRepo.GetImageByImageId(ctx, imageID)

	if err != nil {
		return ErrImageNotFound
	}

	return r.imageRepo.DeleteImageByID(ctx, imageID)
}

func (r *ImageServiceImpl) GetImagesByProductID(ctx context.Context, productID uuid.UUID) ([]*domain.Image, error) {

	if productID == uuid.Nil {
		return nil, ErrInvalidID
	}

	images, err := r.imageRepo.GetImagesByProductID(ctx, productID)

	if err != nil {
		return nil, err
	}

	if images == nil {
		return []*domain.Image{}, nil
	}

	return images, nil
}

func (r *ImageServiceImpl) GetImageByImageId(ctx context.Context, imageID uuid.UUID) (*domain.Image, error) {
	if imageID == uuid.Nil {
		return nil, ErrInvalidID
	}

	image, err := r.imageRepo.GetImageByImageId(ctx, imageID)

	if image == nil {
		return nil, ErrImageNotFound
	}

	if err != nil {
		return nil, err
	}

	return image, nil
}
