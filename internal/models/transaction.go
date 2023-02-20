package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type TransactionType string

const (
	Deposit  = "Deposit"
	Withdraw = "Withdraw"
	Transfer = "Transfer"
)

type Transaction struct {
	gorm.Model
	Type            TransactionType `sql:"type:ENUM('Deposit', 'Withdrawal', 'Transfer')" gorm:"column:transaction_type"`
	Currency        Currency        `sql:"type:ENUM('USD', 'COP', 'MXN')" gorm:"column:currency_type"`
	Amount          decimal.Decimal `gorm:"type:decimal(10,2);"`
	ThisAccount     Account         `gorm:"foreignKey:ID"`
	ReceiverAccount Account         `gorm:"foreignKey:ID"`
}
