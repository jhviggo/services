package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jhviggo/refuels"
	"github.com/jhviggo/repository"
	"github.com/jhviggo/users"
	"github.com/jhviggo/vehicles"
	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Println("[warning] Could not find .env file, using raw env variables.")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "1337"
	}

	if os.Getenv("GIN_MODE") != "release" {
		os.Setenv("RUN_ADDRESS", "localhost")
	}
	fmt.Println(users.HashPasswordWithSalt("123456"))

	repository.DatabaseConnect()
	repository.TestConnection()

	router := gin.Default()
	router.GET("/v1/users", users.GetHandler)
	router.POST("/v1/users", users.PostHandler)
	router.GET("/v1/users/:user/vehicles", vehicles.GetHandler)
	router.POST("/v1/users/:user/vehicles", vehicles.PostHandler)
	router.GET("/v1/users/:user/vehicles/:vehicle/refuels", refuels.GetHandler)
	router.POST("/v1/users/:user/vehicles/:vehicle/refuels", refuels.PostHandler)
	router.GET("/health", healthCheck)
	router.NoRoute(defaultEndpoint)

	fmt.Println("[Server] ⚡ running on 127.0.0.1:" + port)
	router.Run(os.Getenv("RUN_ADDRESS") + ":" + port)
}

func healthCheck(c *gin.Context) {
	repository.TestConnection()
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func defaultEndpoint(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"message": "page not found"})
}
