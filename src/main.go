package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jhviggo/repository"
	"github.com/jhviggo/users"
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

	repository.DatabaseConnect()
	repository.TestConnection()

	router := gin.Default()
	router.GET("/v1/users", users.GetHandler)
	router.POST("/v1/users", users.PostHandler)
	router.GET("/health", healthCheck)

	fmt.Println("[Server] ⚡ running on 127.0.0.1:" + port)
	router.Run(os.Getenv("RUN_ADDRESS") + ":" + port)
}

func healthCheck(c *gin.Context) {
	repository.TestConnection()
	c.JSON(200, gin.H{"status": "OK"})
}
