package repository

import (
	"database/sql"
	"fmt"
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
