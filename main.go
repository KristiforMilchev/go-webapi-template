package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	connectionString := Configuration.GetKey("ConnectionString").(string)
	storage := implementations.Storage{
		ConnectionString: connectionString,
	}

	authMiddlewhere := middlewhere.AuthenticationMiddlewhere{
		JwtService: &jwt,
	}

	authController := &routes.AuthenticationController{
		JwtService: &jwt,
		Storage:    &storage,
	}
	accountsController := &routes.AccountsController{
		Storage: &storage,
	}

	router := gin.New()
	v1 := router.Group("/v1")

	authController.Init(v1)
	accountsController.Init(v1, &authMiddlewhere)

	srv := &http.Server{
		Addr:    port,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	got := <-quit
	fmt.Println(got)
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	connectionDone := <-ctx.Done()
	fmt.Println(connectionDone)
	log.Println("Server exiting")
}
