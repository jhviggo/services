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

	repository.DatabaseConnect()
	repository.TestConnection()

	router := gin.Default()
	router.GET("/v1/users", users.GetUsers)
	router.GET("/health", healthCheck)

	fmt.Println("[Server] ⚡ running on 127.0.0.1:" + port)
	router.Run(":" + port)
}

func healthCheck(c *gin.Context) {
	repository.TestConnection()
	c.JSON(200, gin.H{"status": "OK"})
}
