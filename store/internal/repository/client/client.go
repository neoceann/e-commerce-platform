package repository

import (
	"context"
	"store/internal/domain"
	"store/internal/dto"
	"github.com/google/uuid"
)

type ClientRepository interface {
	CreateClient(ctx context.Context, request *dto.CreateClientRequest) (*domain.Client, error)

	DeleteClient(ctx context.Context, clientID uuid.UUID) error

	GetClientsByFullName(ctx context.Context, request *dto.GetClientsByNameRequest) ([]*domain.Client, error)

	GetClientsWithPagination(ctx context.Context, request *dto.GetclientsWithPaginationRequest) ([]*domain.Client, error)

	GetAllClients(ctx context.Context) ([]*domain.Client, error)

	UpdateClientAddr(ctx context.Context, clientID uuid.UUID, request *dto.UpdateAddressParamsRequest) (*domain.Client, error)

	GetClientByID(ctx context.Context, clientID uuid.UUID) (*domain.Client, error)
}
