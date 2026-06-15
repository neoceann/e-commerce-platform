package domain

import (
	"github.com/google/uuid"
	"time"
)

type Image struct {
	ID        uuid.UUID `json:"id"`
	ImageData []byte    `json:"image_data"`
	ProductID uuid.UUID `json:"product_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
