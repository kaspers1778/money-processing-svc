package repositories

import (
	"fmt"
	"github.com/kaspers1778/money-processing-svc/internal/models"
	"gorm.io/gorm"
)

type AccountRepository interface {
	CreateAccount(account *models.Account)
	GetAccounts(params map[string]interface{}) []*models.Account
	GetAccount(id uint) (*models.Account, error)
	UpdateAccount(id uint, account models.Account) error
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

func (r *AccountRepo) GetAccount(id uint) (*models.Account, error) {
	var account *models.Account
	if err := r.InstanceDB.Where("id = ?", id).First(&account).Error; err != nil {
		return nil, fmt.Errorf("no account with id: %w", id)
	}
	return account, nil
}

func (r *AccountRepo) UpdateAccount(id uint, account models.Account) error {
	var oldAccount *models.Account
	if err := r.InstanceDB.Where("id = ?", id).First(&oldAccount).Error; err != nil {
		return fmt.Errorf("cannot update account: %w", id)
	}
	r.InstanceDB.Model(&oldAccount).Updates(account)
	return nil
}
