package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kaspers1778/money-processing-svc/internal/models"
	"github.com/kaspers1778/money-processing-svc/internal/services"
	"net/http"
	"strconv"
)

type TransactionController struct {
	Service services.TransactionService
}

func NewTransactionController(service services.TransactionService) *TransactionController {
	return &TransactionController{service}
}

func (cc *TransactionController) CreateTransaction(c *gin.Context) {
	var input models.TransactionRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := cc.Service.CreateTransaction(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": input})
}

func (cc *TransactionController) GetTransactionsByAccount(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong data."})
	}
	transactions, err := cc.Service.GetTransactionByAccount(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"data": transactions})
}
