package users

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
}

var users = []User{
	{ID: "1", Name: "Viggo", Username: "jhviggo"},
	{ID: "2", Name: "Jimmy", Username: "jimmyjohn"},
	{ID: "3", Name: "Sarah", Username: "thesarah123"},
}

/* Handlers */
func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func PostUser(c *gin.Context) {
	var newUser User

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
		errorMessage := fmt.Sprintf("Unable to create user. Missing values [%s]", strings.Join(missingValues[:], ", "))
		log.Println(errorMessage)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": errorMessage})
		return
	}

	c.IndentedJSON(http.StatusOK, users[0])
}

/* Logic */
func HashPasswordWithSalt(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err == nil {
		log.Println("Unable to generate hash of password with length", len(password))
		return "", err
	}

	return string(bytes), nil
}

func ValidatePassword(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
