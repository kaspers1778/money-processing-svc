package repositories

import (
	"github.com/kaspers1778/money-processing-svc/internal/models"
	"gorm.io/gorm"
)

type AccountRepository interface {
	CreateAccount(account *models.Account)
	GetAccounts(params map[string]interface{}) []*models.Account
}

type AccountRepo struct {
	InstanceDB *gorm.DB
}

func NewAccountRepo(instanceDB *gorm.DB) AccountRepository {
	return &AccountRepo{instanceDB}
}

func (r *AccountRepo) CreateAccount(account *models.Account) {
	r.InstanceDB.Create(account)
}

func (r *AccountRepo) GetAccounts(params map[string]interface{}) []*models.Account {
	var accounts []*models.Account
	r.InstanceDB.Where(params).Find(&accounts)
	return accounts
}
