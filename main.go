package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	implementations "leanmeal/api/Implementations"
	"leanmeal/api/interfaces"
	"leanmeal/api/middlewhere"
	"leanmeal/api/routes"
)

func main() {

	config := implementations.Configuration{}
	config.Load()

	startServer(&config)
}

func startServer(Configuration interfaces.Configuration) {
	port := Configuration.GetKey("Port").(string)

	jwt := implementations.JwtService{}

	jwt.Secret = Configuration.GetKey("jwt-key").(string)
	jwt.Issuer = Configuration.GetKey("jwt-issuer").(string)

	var customers implementations.CustomerService
	TestService(&customers)
	connectionString := Configuration.GetKey("ConnectionString").(string)

	authController := &routes.AuthenticationController{
		JwtService: &jwt,
		Storage: &implementations.Storage{
			ConnectionString: connectionString,
		},
	}

	router := gin.New()
	router.Use(middlewhere.Authorize())

	authController.InitAuthenticationRouter(router)

	router.Run(port)
}

func TestService(customers interfaces.ICustomerService) {
	customers.Add("Kristifor", 22)
	customers.Add("Ivan", 22)

	data := customers.Get()

	for d := range data {
		fmt.Println(d)
	}
}
