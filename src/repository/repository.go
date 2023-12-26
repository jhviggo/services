package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func DatabaseConnect() {
	databaseSource := os.Getenv("DATABASE_SOURCE")
	databaseName := os.Getenv("DATABASE_NAME")
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	databaseUsername := os.Getenv("DATABASE_USERNAME")

	if databaseSource == "" {
		panic("Missing ENV variable [DATABASE_SOURCE]")
	} else if databaseName == "" {
		panic("Missing ENV variable [DATABASE_NAME]")
	} else if databasePassword == "" {
		panic("Missing ENV variable [DATABASE_PASSWORD]")
	} else if databaseUsername == "" {
		panic("Missing ENV variable [DATABASE_USERNAME]")
	}

	var err error
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", databaseUsername, databasePassword, databaseSource, databaseName))

	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}

func TestConnection() {
	fmt.Print("Testing connection to database...")
	pingErr := db.Ping()
	rows, err := db.Query("SELECT 1+1 AS number")
	if pingErr != nil || err != nil {
		panic("\nCould not connect to Database!")
	}
	defer rows.Close()

	i := 0
	for rows.Next() {
		rows.Scan(&i)
	}

	if i != 2 {
		panic("\nReturned value from database was wrong!")
	}
	fmt.Println("OK!")
}

func AddUser(user User) (User, error) {
	var addedUser User

	_, insertErr := db.Exec("INSERT INTO users (displayName, username, passwd, createdAt) VALUES (?,?,?,NOW());",
		user.Name, user.Username, user.Password)

	if insertErr != nil {
		log.Println("unable to insert user into database")
		return addedUser, insertErr
	}
	selectErr := db.QueryRow("SELECT id, displayName, username FROM users WHERE username = ?;", user.Username).Scan(&addedUser.ID, &addedUser.Name, &addedUser.Username)

	if selectErr != nil {
		log.Println("unable to get user from database")
		return addedUser, selectErr
	}
	return addedUser, nil
}

func GetUsers() ([]User, error) {
	var users []User

	rows, err := db.Query("SELECT id, displayName, username FROM users;")
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
