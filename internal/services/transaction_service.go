package services

import (
	"fmt"
	"github.com/kaspers1778/money-processing-svc/internal/models"
	"github.com/kaspers1778/money-processing-svc/internal/repositories"
	"github.com/shopspring/decimal"
)

type TransactionService interface {
	CreateTransaction(transaction models.TransactionRequest) error
	GetTransactionByAccount(accountId uint) ([]models.Transaction, error)
}

type TransactionSrc struct {
	repository repositories.TransactionRepository
	accountSrc AccountService
}

func NewTransactionsSrc(repo repositories.TransactionRepository, src AccountService) TransactionService {
	return &TransactionSrc{repo, src}
}

func (s *TransactionSrc) CreateTransaction(transaction models.TransactionRequest) error {
	switch transaction.Type {
	case models.Deposit, models.Withdraw, models.Transfer:
		break
	default:
		return fmt.Errorf("unavailable operation: %w", transaction.Type)
	}

	switch transaction.Currency {
	case models.USD, models.COP, models.MXN:
		break
	default:
		return fmt.Errorf("unavailable currency: %w", transaction.Currency)
	}

	account, err := s.accountSrc.GetAccountByID(transaction.ThisAccount)
	if err != nil {
		return err
	}
	if account.Currency != transaction.Currency {
		return fmt.Errorf("account currency %w is differrent from transaction currency %w", account.Currency,
			transaction.Currency)
	}

	switch transaction.Type {
	case models.Deposit:
		if err = s.accountSrc.UpdateAccountAmount(account.ID, decimal.Sum(account.Amount, transaction.Amount)); err != nil {
			return fmt.Errorf("cannot deposit to account:%w", err)
		}
	case models.Withdraw:
		newAmount := account.Amount.Sub(transaction.Amount)
		if newAmount.IsNegative() {
			return fmt.Errorf("not enough money:%w", err)
		}
		if err = s.accountSrc.UpdateAccountAmount(account.ID, newAmount); err != nil {
			return fmt.Errorf("cannot withdraw from account:%w", err)
		}
	case models.Transfer:
		receiverAccount, err := s.accountSrc.GetAccountByID(transaction.ThisAccount)
		if err != nil {
			return err
		}
		if receiverAccount.Currency != transaction.Currency {
			return fmt.Errorf("account currency %w is differrent from transaction currency %w", account.Currency,
				transaction.Currency)
		}
		newAmount := account.Amount.Sub(transaction.Amount)
		if newAmount.IsNegative() {
			return fmt.Errorf("not enough money:%w", err)
		}
		if err = s.accountSrc.UpdateAccountAmount(account.ID, newAmount); err != nil {
			return fmt.Errorf("cannot transfer from account:%w", err)
		}

		if err = s.accountSrc.UpdateAccountAmount(receiverAccount.ID, decimal.Sum(receiverAccount.Amount, transaction.Amount)); err != nil {
			return fmt.Errorf("cannot transfer to account:%w", err)
		}
	default:
		return fmt.Errorf("unavailable operation: %w", transaction.Type)
	}
	return nil
}

func (s *TransactionSrc) GetTransactionByAccount(accountId uint) ([]models.Transaction, error) {
	return s.repository.GetTransactionByAccount(accountId)
}
