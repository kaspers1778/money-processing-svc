package repositories

import (
	"github.com/kaspers1778/money-processing-svc/internal/models"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction *models.Transaction)
	GetTransactionByAccount(accountID int) []*models.Transaction
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

func (r *TransactionRepo) GetTransactionByAccount(accountID int) []*models.Transaction {
	var transactions []*models.Transaction
	r.InstanceDB.Where("account = ?", accountID).Find(&transactions)
	return transactions
}
