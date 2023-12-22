package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jhviggo/users"
)

func main() {
	fmt.Println("main go!")
	router := gin.Default()
	router.GET("/users", users.GetUsers)
	router.Run("localhost:1337")
}