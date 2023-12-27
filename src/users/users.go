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

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
}

/* Handlers */
func GetHandler(c *gin.Context) {
	users, err := fetchUsers()
	if err != nil {
		log.Println("unable to get users", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "unable to get users"})
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}

func PostHandler(c *gin.Context) {
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
		errorMessage := fmt.Sprintf("unable to create user. Missing values [%s]", strings.Join(missingValues[:], ", "))
		log.Println(errorMessage)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": errorMessage})
		return
	}

	user, err := insertUser(newUser)
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

	if err != nil {
		log.Println("unable to generate hash of password with length", len(password))
		return "", err
	}
	return string(bytes), nil
}

func ValidatePassword(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

/* Database queries */
func insertUser(user User) (User, error) {
	var addedUser User
	hashedPassword, hashErr := HashPasswordWithSalt(user.Password)
	fmt.Println("password:", hashedPassword)

	if hashErr != nil {
		log.Println("failed to hash user's password")
		return addedUser, hashErr
	}

	_, insertErr := repository.DB.Exec("INSERT INTO users (displayName, username, passwd, createdAt) VALUES (?,?,?,NOW());",
		user.Name, user.Username, hashedPassword)

	if insertErr != nil {
		log.Println("unable to insert user into database")
		return addedUser, insertErr
	}
	selectErr := repository.DB.QueryRow("SELECT id, displayName, username FROM users WHERE username = ?;", user.Username).Scan(&addedUser.ID, &addedUser.Name, &addedUser.Username)

	if selectErr != nil {
		log.Println("unable to get user from database")
		return addedUser, selectErr
	}
	return addedUser, nil
}

func fetchUsers() ([]User, error) {
	var users []User

	rows, err := repository.DB.Query("SELECT id, displayName, username FROM users;")
	if err != nil {
		return make([]User, 0), err
	}

	for rows.Next() {
		var myUser User
		rows.Scan(&myUser.ID, &myUser.Name, &myUser.Username)
		users = append(users, myUser)
	}

	return users, nil
}
