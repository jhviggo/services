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
	port := os.Getenv("PORT")
	if port == "" {
		port = "1337"
	}

	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Println("[warning] Could not find .env file, using raw env variables.")
	}

	repository.DatabaseConnect()
	repository.TestConnection()

	router := gin.Default()
	router.GET("/v1/users", users.GetUsers)
	router.GET("/health", healthCheck)

	fmt.Println("[Server] ⚡ running on localhost:" + port)
	router.Run("localhost:" + port)
}

func healthCheck(c *gin.Context) {
	repository.TestConnection()
	c.JSON(200, gin.H{"status": "OK"})
}
