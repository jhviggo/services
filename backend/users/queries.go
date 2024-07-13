package users

import (
	"fmt"
	"log"

	"viggo.work/repository"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Token    string `json:"token,omitempty"`
}

func insertUser(user User) (User, error) {
	var addedUser User
	hashedPassword, err := HashPasswordWithSalt(user.Password)

	if err != nil {
		log.Println("[warning] failed to hash user's password")
		return addedUser, err
	}

	_, err = repository.DB.Exec(
		"INSERT INTO users (displayName, username, passwd, createdAt) VALUES (?,?,?,NOW());",
		user.Name,
		user.Username,
		hashedPassword)

	if err != nil {
		log.Println("[warning] unable to insert user into database")
		return addedUser, err
	}
	err = repository.DB.QueryRow(
		"SELECT id, displayName, username FROM users WHERE username = ?;",
		user.Username).Scan(&addedUser.ID, &addedUser.Name, &addedUser.Username)

	if err != nil {
		log.Println("[warning] unable to get user from database")
		return addedUser, err
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

func fetchValidUser(username string, password string) (User, error) {
	var user User
	var userPassword string

	err := repository.DB.QueryRow("SELECT id, username, displayName, passwd FROM users WHERE username = ? LIMIT 1;", username).Scan(&user.ID, &user.Username, &user.Name, &userPassword)
	if err != nil {
		return user, fmt.Errorf("user does not exists")
	}

	if err = ValidatePassword(userPassword, password); err != nil {
		return user, fmt.Errorf("invalid password")
	}
	return user, nil
}
