package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kaspers1778/money-processing-svc/internal/controllers"
	"github.com/kaspers1778/money-processing-svc/internal/db"
	"github.com/kaspers1778/money-processing-svc/internal/repositories"
	"github.com/kaspers1778/money-processing-svc/internal/services"
	"go.uber.org/dig"
	"log"
)

type API struct {
	ClientController      *controllers.ClientController
	AccountController     *controllers.AccountController
	TransactionController *controllers.TransactionController
}

func NewAPI(clientController *controllers.ClientController, accountController *controllers.AccountController, transactionController *controllers.TransactionController) *API {
	return &API{ClientController: clientController, AccountController: accountController, TransactionController: transactionController}
}

func main() {
	c := dig.New()
	c.Provide(db.ConnectDB)
	c.Provide(repositories.NewClientRepo)
	c.Provide(repositories.NewAccountRepo)
	c.Provide(repositories.NewTransactionRepo)
	c.Provide(services.NewClientSrc)
	c.Provide(services.NewAccountSrc)
	c.Provide(services.NewTransactionsSrc)
	c.Provide(controllers.NewClientController)
	c.Provide(controllers.NewAccountController)
	c.Provide(controllers.NewTransactionController)
	c.Provide(NewAPI)
	err := c.Invoke(func(api *API) {
		r := gin.Default()
		r.GET("/clients", api.ClientController.GetClients)
		r.GET("/clients/:email", api.ClientController.GetClientByEmail)
		r.POST("/clients", api.ClientController.CreateClient)
		r.GET("/accounts", api.AccountController.GetAccounts)
		r.POST("/accounts", api.AccountController.CreateAccount)
		r.GET("/transactions/:id", api.TransactionController.GetTransactionsByAccount)
		r.POST("/transactions", api.TransactionController.CreateTransaction)
		r.Run()
	})
	if err != nil {
		log.Fatal(err)
	}

}
