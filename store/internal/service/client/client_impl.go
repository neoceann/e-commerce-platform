package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"store/internal/domain"
	"store/internal/dto"
	"store/internal/repository/client"
	"strings"
)

type ClientServiceImpl struct {
	clientRepo repository.ClientRepository
}

func NewClientService(repo repository.ClientRepository) ClientService {
	return &ClientServiceImpl{clientRepo: repo}
}

func (s *ClientServiceImpl) CreateClient(ctx context.Context, request *dto.CreateClientRequest) (*domain.Client, error) {
	if strings.TrimSpace(request.ClientName) == "" {
		return nil, fmt.Errorf("%w: client name is required", ErrInvalidClientData)
	}
	if strings.TrimSpace(request.ClientSurname) == "" {
		return nil, fmt.Errorf("%w: client surname is required", ErrInvalidClientData)
	}

	client, err := s.clientRepo.CreateClient(ctx, request)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (s *ClientServiceImpl) DeleteClient(ctx context.Context, clientID uuid.UUID) error {
	if clientID == uuid.Nil {
		return ErrInvalidID
	}

	_, err := s.clientRepo.GetClientByID(ctx, clientID)

	if err != nil {
		return ErrClientNotFound
	}

	return s.clientRepo.DeleteClient(ctx, clientID)
}

func (s *ClientServiceImpl) GetClientsByFullName(ctx context.Context, request *dto.GetClientsByNameRequest) ([]*domain.Client, error) {
	if strings.TrimSpace(request.Name) == "" {
		return nil, fmt.Errorf("%w: client name is required", ErrInvalidClientData)
	}

	if strings.TrimSpace(request.SurName) == "" {
		return nil, fmt.Errorf("%w: client surname is required", ErrInvalidClientData)
	}

	clients, err := s.clientRepo.GetClientsByFullName(ctx, request)

	if err != nil {
		return nil, err
	}

	if clients == nil {
		return []*domain.Client{}, nil
	}

	return clients, nil
}

func (s *ClientServiceImpl) GetClientsWithPagination(ctx context.Context, request *dto.GetclientsWithPaginationRequest) ([]*domain.Client, error) {

	if request.Limit == nil {
		return s.clientRepo.GetAllClients(ctx)
	}

	if *request.Limit <= 0 || *request.Limit > 100 {
		return nil, ErrInvalidPagination
	}

	if request.Offset != nil {
		if *request.Offset < 0 {
			return nil, ErrInvalidPagination
		}
	} else {
		var zeroOffset int32 = 0
		data := &dto.GetclientsWithPaginationRequest{Limit: request.Limit, Offset: &zeroOffset}
		return s.clientRepo.GetClientsWithPagination(ctx, data)
	}

	clients, err := s.clientRepo.GetClientsWithPagination(ctx, request)

	if err != nil {
		return nil, err
	}

	if clients == nil {
		return []*domain.Client{}, nil
	}

	return clients, nil
}

func (s *ClientServiceImpl) UpdateClientAddr(ctx context.Context, clientID uuid.UUID, request *dto.UpdateAddressParamsRequest) (*domain.Client, error) {
	if clientID == uuid.Nil {
		return nil, ErrInvalidID
	}

	if strings.TrimSpace(request.Country) == "" ||
		strings.TrimSpace(request.City) == "" ||
		strings.TrimSpace(request.Street) == "" {
		return nil, ErrInvalidAddrData
	}

	_, err := s.clientRepo.GetClientByID(ctx, clientID)

	if err != nil {
		return nil, ErrClientNotFound
	}

	return s.clientRepo.UpdateClientAddr(ctx, clientID, request)
}
