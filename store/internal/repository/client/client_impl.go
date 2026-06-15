package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"store/internal/domain"
	"store/internal/dto"
	"store/internal/repository/db"
)

type ClientRepositoryImpl struct {
	queries *db.Queries
}

func NewClientRepository(q *db.Queries) ClientRepository {
	return &ClientRepositoryImpl{queries: q}
}

func (r *ClientRepositoryImpl) CreateClient(ctx context.Context, request *dto.CreateClientRequest) (*domain.Client, error) {

	client, err := r.queries.CreateClient(ctx, db.CreateClientParams{
		ClientName:    request.ClientName,
		ClientSurname: request.ClientSurname,
		Birthday:      request.Birthday,
		Gender:        request.Gender,
		AddressID:     request.AddressID})

	if err != nil {
		return nil, fmt.Errorf("failed to create new client: %w", err)
	}

	return dto.ClientFromDbToDomain(client), nil
}

func (r *ClientRepositoryImpl) DeleteClient(ctx context.Context, clientID uuid.UUID) error {
	return r.queries.DeleteClient(ctx, clientID)
}

func (r *ClientRepositoryImpl) GetClientsByFullName(ctx context.Context, request *dto.GetClientsByNameRequest) ([]*domain.Client, error) {
	clients, err := r.queries.GetClientsByFullName(ctx, db.GetClientsByFullNameParams{ClientName: request.Name, ClientSurname: request.SurName})

	if err != nil {
		return nil, fmt.Errorf("failed to get clients: %w", err)
	}

	var domainClients []*domain.Client

	for _, client := range clients {
		domainClients = append(domainClients, dto.ClientFromDbToDomain(client))
	}

	return domainClients, nil
}

func (r *ClientRepositoryImpl) GetClientsWithPagination(ctx context.Context, request *dto.GetclientsWithPaginationRequest) ([]*domain.Client, error) {
	clients, err := r.queries.GetClientsWithPagination(ctx, db.GetClientsWithPaginationParams{Limit: *request.Limit, Offset: *request.Offset})

	if err != nil {
		return nil, fmt.Errorf("failed to get clients: %w", err)
	}

	var domainClients []*domain.Client

	for _, client := range clients {
		domainClients = append(domainClients, dto.ClientFromDbToDomain(client))
	}

	return domainClients, nil
}

func (r *ClientRepositoryImpl) GetAllClients(ctx context.Context) ([]*domain.Client, error) {
	clients, err := r.queries.GetAllClients(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to get clients: %w", err)
	}

	var domainClients []*domain.Client

	for _, client := range clients {
		domainClients = append(domainClients, dto.ClientFromDbToDomain(client))
	}

	return domainClients, nil
}

func (r *ClientRepositoryImpl) UpdateClientAddr(ctx context.Context, clientID uuid.UUID, request *dto.UpdateAddressParamsRequest) (*domain.Client, error) {

	newAddr, err := r.queries.CreateAddress(ctx, db.CreateAddressParams{Country: request.Country,
		City:   request.City,
		Street: request.Street})

	if err != nil {
		return nil, fmt.Errorf("failed to create address: %w", err)
	}

	client, err := r.queries.UpdateClientAddress(ctx, db.UpdateClientAddressParams{ID: clientID, AddressID: newAddr.ID})

	if err != nil {
		return nil, fmt.Errorf("failed to update client address: %w", err)
	}

	return dto.ClientFromDbToDomain(client), nil
}

func (r *ClientRepositoryImpl) GetClientByID(ctx context.Context, clientID uuid.UUID) (*domain.Client, error) {

	client, err := r.queries.GetClientByID(ctx, clientID)

	if err != nil {
		return nil, fmt.Errorf("client not found")
	}

	return dto.ClientFromDbToDomain(client), nil

}
