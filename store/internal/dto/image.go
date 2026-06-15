package dto

import (
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"store/internal/domain"
	"store/internal/repository/db"
	"time"
)

type ImageResponse struct {
	ID        uuid.UUID `json:"id"`
	ProductID uuid.UUID `json:"product_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ImageFromDbToDomain(i db.Image) *domain.Image {
	return &domain.Image{
		ID:        i.ID,
		ImageData: i.ImageData,
		ProductID: i.ProductID,
		CreatedAt: i.CreatedAt,
		UpdatedAt: i.UpdatedAt,
	}
}

func ImageToResponce(i *domain.Image) *ImageResponse {
	return &ImageResponse{
		ID:        i.ID,
		ProductID: i.ProductID,
		CreatedAt: i.CreatedAt,
		UpdatedAt: i.UpdatedAt,
	}
}

type CreateImageRequest struct {
	ImageData string     `json:"image_data" validate:"required"`
	ProductID *uuid.UUID `json:"product_id" validate:"required"`
}

type CreateImageDataByte struct {
	ImageData []byte
	ProductID *uuid.UUID
}

type UpdateImageRequest struct {
	ImageData string `json:"image_data" validate:"required"`
}
type UpdateImageDataByte struct {
	ImageData []byte
}

func StringToBytes(imageData string) ([]byte, error) {
	if imageData == "" {
		return nil, fmt.Errorf("imageData required")
	}

	data, err := base64.StdEncoding.DecodeString(imageData)
	if err != nil {
		data, err = base64.URLEncoding.DecodeString(imageData)
		if err != nil {
			return nil, fmt.Errorf("invalid base64 encoding: %w", err)
		}
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("image data is empty")
	}

	return data, nil
}
