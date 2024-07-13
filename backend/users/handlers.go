package users

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetHandler(c *gin.Context) {
	users, err := fetchUsers()
	if err != nil {
		log.Println("[warning] unable to get users", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "unable to get users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func PostHandler(c *gin.Context) {
	var newUser User

	if err := c.BindJSON(&newUser); err != nil {
		log.Printf("[warning] unable to create user. %s\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	var missingValues []string

	if newUser.Username == "" {
		missingValues = append(missingValues, "username")
	}
	if newUser.Password == "" {
		missingValues = append(missingValues, "password")
	}
	if newUser.Name == "" {
		missingValues = append(missingValues, "name")
	}

	if len(missingValues) > 0 {
		errorMessage := fmt.Sprintf("unable to create user. Missing values [%s]", strings.Join(missingValues[:], ", "))
		log.Println(errorMessage)
		c.JSON(http.StatusBadRequest, gin.H{"message": errorMessage})
		return
	}

	user, err := insertUser(newUser)
	if err != nil {
		log.Println("[warning] failed to add user to database", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to add user to database"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func Login(c *gin.Context) {
	var userAuth User

	if err := c.BindJSON(&userAuth); err != nil {
		fmt.Println("User:", userAuth.Username, userAuth.Password)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Missing username or password"})
		return
	}

	user, err := fetchValidUser(userAuth.Username, userAuth.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid username or password"})
		return
	}

	token, err := CreateToken(userAuth.Username)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create token"})
		return
	}

	user.Token = token
	c.JSON(http.StatusOK, user)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/v1/login" || c.Request.URL.Path == "/health" {
			return
		}

		authHeader := strings.Replace(c.GetHeader("Authorization"), "Bearer ", "", 1)
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Missing Authorization header"})
			c.Abort()
			return
		}

		err := VerifyToken(authHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
		}
	}
}
