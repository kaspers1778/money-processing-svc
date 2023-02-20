package services

import (
	"fmt"
	"github.com/kaspers1778/money-processing-svc/internal/models"
	"github.com/kaspers1778/money-processing-svc/internal/pkg"
	"github.com/kaspers1778/money-processing-svc/internal/repositories"
	"net/url"
)

type ClientService interface {
	CreateClient(client models.ClientRequest) error
	GetClients(params url.Values) []*models.Client
	GetClientByEmail(client models.ClientRequest) (*models.Client, error)
}

type ClientSrc struct {
	repository repositories.ClientRepository
}

func NewClientSrc(repo repositories.ClientRepository) ClientService {
	return &ClientSrc{repo}
}

func (s *ClientSrc) CreateClient(client models.ClientRequest) error {
	if s.repository.IsClientExists(client.Email) {
		return fmt.Errorf("cannot create user with email: %w", client.Email)
	}
	s.repository.CreateClient(&models.Client{
		Email:    client.Email,
		Accounts: nil,
	})
	return nil
}

func (s *ClientSrc) GetClients(params url.Values) []*models.Client {
	return s.repository.GetClients(pkg.ParseQueryParams(params))
}

func (s *ClientSrc) GetClientByEmail(client models.ClientRequest) (*models.Client, error) {
	return s.repository.GetClientByEmail(client.Email)
}
