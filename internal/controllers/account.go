package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kaspers1778/money-processing-svc/internal/models"
	"github.com/kaspers1778/money-processing-svc/internal/services"
	"net/http"
)

type AccountController struct {
	Service services.AccountService
}

func NewAccountController(service services.AccountService) *AccountController {
	return &AccountController{service}
}

func (cc *AccountController) GetAccounts(c *gin.Context) {
	params := c.Request.URL.Query()
	accounts := cc.Service.GetAccounts(params)
	c.JSON(http.StatusOK, gin.H{"data": accounts})
}

func (cc *AccountController) CreateAccount(c *gin.Context) {
	var input models.AccountRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := cc.Service.CreateAccount(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": input})
}
