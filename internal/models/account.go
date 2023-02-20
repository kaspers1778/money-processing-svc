package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Currency string

const (
	USD Currency = "USD"
	COP Currency = "COP"
	MXN Currency = "MXN"
)

type Account struct {
	gorm.Model
	Owner    uint
	Currency Currency        `sql:"type:ENUM('USD', 'COP', 'MXN')" gorm:"column:currency_type"`
	Amount   decimal.Decimal `gorm:"type:decimal(10,2);"`
}

type AccountRequest struct {
	Owner    uint            `json:"owner" binding:"required"`
	Currency Currency        `json:"currency" binding:"required"`
	Amount   decimal.Decimal `json:"amount" binding:"required"`
}
