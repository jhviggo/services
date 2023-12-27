package users

import (
	"log"

	"github.com/jhviggo/repository"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
}

func insertUser(user User) (User, error) {
	var addedUser User
	hashedPassword, hashErr := HashPasswordWithSalt(user.Password)

	if hashErr != nil {
		log.Println("[warning] failed to hash user's password")
		return addedUser, hashErr
	}

	_, insertErr := repository.DB.Exec("INSERT INTO users (displayName, username, passwd, createdAt) VALUES (?,?,?,NOW());",
		user.Name, user.Username, hashedPassword)

	if insertErr != nil {
		log.Println("[warning] unable to insert user into database")
		return addedUser, insertErr
	}
	selectErr := repository.DB.QueryRow("SELECT id, displayName, username FROM users WHERE username = ?;", user.Username).Scan(&addedUser.ID, &addedUser.Name, &addedUser.Username)

	if selectErr != nil {
		log.Println("[warning] unable to get user from database")
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
