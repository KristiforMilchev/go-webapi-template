package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	implementations "leanmeal/api/Implementations"
	"leanmeal/api/interfaces"
)

type AuthenticationController struct {
	AuthenticationService interfaces.AuthenticationService
}

func (authController *AuthenticationController) beginRequest(ctx *gin.Context) {
	email := ctx.Param("email")

	if email == "" {
		ctx.JSON(500, "Bad request")
	}

	fmt.Println(email)
	data := authController.AuthenticationService.GetMessage(email)
	ctx.JSON(http.StatusOK, gin.H{"message": data})
}

func (authController *AuthenticationController) InitAuthenticationRouter(r *gin.Engine) {
	authController.AuthenticationService = &implementations.AuthenticationService{}
	println("initializing Authentication Controller")
	go authController.AuthenticationService.Start()
	// Router group
	v1 := r.Group("/v1")
	{
		v1.GET("/begin-request/:email", authController.beginRequest)
	}

}
