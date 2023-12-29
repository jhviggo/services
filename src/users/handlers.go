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
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unable to unmarshal JSON"})
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
