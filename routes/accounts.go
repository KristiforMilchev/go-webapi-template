package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"leanmeal/api/interfaces"
	"leanmeal/api/middlewhere"
	"leanmeal/api/repositories"
)

type AccountsController struct {
	AccountRepository repositories.Accounts
	Storage           interfaces.Storage
}

func (ac *AccountsController) getAccountById(ctx *gin.Context) {
	email := ctx.Param("email")

	if email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Missing parameter"})
	}

	ac.Storage.Open(ac.AccountRepository.ConnectionString)
	account := ac.AccountRepository.UserExists(email)
	ac.Storage.Close()

	ctx.JSON(http.StatusOK, account)
}

func (ac *AccountsController) Init(r *gin.RouterGroup, authMiddlewhere *middlewhere.AuthenticationMiddlewhere) {
	accounts := r.Group("accounts")

	accounts.Use(authMiddlewhere.Authorize())
	accounts.GET(":id", ac.getAccountById)
}
