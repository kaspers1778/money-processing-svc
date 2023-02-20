package services

import (
	"fmt"
	"github.com/kaspers1778/money-processing-svc/internal/models"
	"github.com/kaspers1778/money-processing-svc/internal/pkg"
	"github.com/kaspers1778/money-processing-svc/internal/repositories"
	"github.com/shopspring/decimal"
	"net/url"
)

type AccountService interface {
	CreateAccount(account models.AccountRequest) error
	GetAccounts(params url.Values) []*models.Account
	GetAccountByID(accountID uint) (*models.Account, error)
	GetAccountCurrencyByID(accountID uint) (models.Currency, error)
	UpdateAccountAmount(id uint, newAmount decimal.Decimal) error
}

type AccountSrc struct {
	repository repositories.AccountRepository
}

func NewAccountSrc(repo repositories.AccountRepository) AccountService {
	return &AccountSrc{repo}
}

func (s *AccountSrc) CreateAccount(account models.AccountRequest) error {
	switch account.Currency {
	case models.USD, models.COP, models.MXN:
		break
	default:
		return fmt.Errorf("unavailable currency: %w", account.Currency)
	}
	if account.Amount.IsNegative() {
		return fmt.Errorf("wring amount: %w", account.Amount)
	}

	s.repository.CreateAccount(&models.Account{
		Owner:    account.Owner,
		Currency: account.Currency,
		Amount:   account.Amount,
	})
	return nil
}

func (s *AccountSrc) GetAccounts(params url.Values) []*models.Account {
	return s.repository.GetAccounts(pkg.ParseQueryParams(params))
}

func (s *AccountSrc) GetAccountByID(accountID uint) (*models.Account, error) {
	return s.repository.GetAccount(accountID)
}

func (s *AccountSrc) GetAccountCurrencyByID(accountID uint) (models.Currency, error) {
	account, err := s.repository.GetAccount(accountID)
	if err != nil {
		return "", err
	}
	return account.Currency, nil
}

func (s *AccountSrc) UpdateAccountAmount(id uint, newAmount decimal.Decimal) error {
	updatedAccount := models.Account{
		Amount: newAmount,
	}
	return s.repository.UpdateAccount(uint(id), updatedAccount)
}
