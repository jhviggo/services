package users

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jhviggo/repository"
	"golang.org/x/crypto/bcrypt"
)

/* Handlers */
func GetUsers(c *gin.Context) {
	users, err := repository.GetUsers()
	if err != nil {
		log.Println("unable to get users", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "unable to get users"})
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}

func PostUser(c *gin.Context) {
	var newUser repository.User

	if err := c.BindJSON(&newUser); err != nil {
		log.Printf("unable to create user. %s\n", err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Unable to unmarshal JSON"})
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
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": errorMessage})
		return
	}

	user, err := repository.AddUser(newUser)
	if err != nil {
		log.Println("failed to add user to database", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Unable to add user to database"})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

/* Logic */
func HashPasswordWithSalt(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err == nil {
		log.Println("unable to generate hash of password with length", len(password))
		return "", err
	}

	return string(bytes), nil
}

func ValidatePassword(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
