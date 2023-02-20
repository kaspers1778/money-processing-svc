package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kaspers1778/money-processing-svc/internal/models"
	"github.com/kaspers1778/money-processing-svc/internal/services"
	"net/http"
)

type ClientController struct {
	Service services.ClientService
}

func NewClientController(service services.ClientService) *ClientController {
	return &ClientController{service}
}

func (cc *ClientController) GetClients(c *gin.Context) {
	params := c.Request.URL.Query()
	clients := cc.Service.GetClients(params)
	c.JSON(http.StatusOK, gin.H{"data": clients})
}

func (cc *ClientController) GetClientByEmail(c *gin.Context) {
	input := c.Param("email")
	client, err := cc.Service.GetClientByEmail(models.ClientRequest{Email: input})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": client})
}

func (cc *ClientController) CreateClient(c *gin.Context) {
	var input models.ClientRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := cc.Service.CreateClient(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": input})
}
