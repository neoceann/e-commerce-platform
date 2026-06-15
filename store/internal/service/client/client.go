package service

import (
	"context"
	"github.com/google/uuid"
	"store/internal/domain"
	"store/internal/dto"
)

type ClientService interface {
	CreateClient(ctx context.Context, request *dto.CreateClientRequest) (*domain.Client, error)

	DeleteClient(ctx context.Context, clientID uuid.UUID) error

	GetClientsByFullName(ctx context.Context, request *dto.GetClientsByNameRequest) ([]*domain.Client, error)

	GetClientsWithPagination(ctx context.Context, request *dto.GetclientsWithPaginationRequest) ([]*domain.Client, error)

	UpdateClientAddr(ctx context.Context, clientID uuid.UUID, request *dto.UpdateAddressParamsRequest) (*domain.Client, error)
}
