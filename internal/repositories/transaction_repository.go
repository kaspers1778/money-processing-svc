package repositories

import (
	"fmt"
	"github.com/kaspers1778/money-processing-svc/internal/models"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction *models.Transaction)
	GetTransactionByAccount(accountID uint) ([]models.Transaction, error)
}

type TransactionRepo struct {
	InstanceDB *gorm.DB
}

func NewTransactionRepo(instanceDB *gorm.DB) TransactionRepository {
	return &TransactionRepo{instanceDB}
}

func (r *TransactionRepo) CreateTransaction(transaction *models.Transaction) {
	r.InstanceDB.Create(transaction)
}

func (r *TransactionRepo) GetTransactionByAccount(accountID uint) ([]models.Transaction, error) {
	var transactions []models.Transaction
	if err := r.InstanceDB.Where("this_account = ?", accountID).Find(&transactions).Error; err != nil {
		return nil, fmt.Errorf("cannot find transacctions by account:%w", err)
	}
	return transactions, nil
}
