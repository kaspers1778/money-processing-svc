package repositories

import (
	"fmt"
	"github.com/kaspers1778/money-processing-svc/internal/models"
	"gorm.io/gorm"
)

type ClientRepository interface {
	CreateClient(client *models.Client)
	GetClientByEmail(email string) (*models.Client, error)
	IsClientExists(email string) bool
}

type ClientRepo struct {
	InstanceDB *gorm.DB
}

func NewClientRepo(instanceDB *gorm.DB) ClientRepository {
	return &ClientRepo{instanceDB}
}

func (r *ClientRepo) CreateClient(client *models.Client) {
	r.InstanceDB.Create(client)
}

func (r *ClientRepo) GetClientByEmail(email string) (*models.Client, error) {
	var client *models.Client
	if err := r.InstanceDB.Preload("accounts").Where("email = ?", email).First(&client).Error; err != nil {
		return nil, fmt.Errorf("cannot find client by email: %w", err)
	}
	return client, nil
}

func (r *ClientRepo) IsClientExists(email string) bool {
	if err := r.InstanceDB.Where("email = ?", email).Error; err != nil {
		return false
	}
	return true
}
