package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Email    string    `json:"email" gorm:"uniqueIndex"`
	Accounts []Account `gorm:"foreignKey:Owner"`
}

type ClientRequest struct {
	Email string `json:"email" binding:"required,email"`
}
