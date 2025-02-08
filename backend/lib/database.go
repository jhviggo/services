package lib

import (
	"database/sql"
	"errors"
	"fmt"
	"services/project/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectToSQLite(dataPath string) {
	fmt.Println("Connecting to DB...💾")
	sqlDB, err := sql.Open("sqlite", dataPath+"/db.db")

	if err != nil {
		panic("failed to connect database")
	}
	db, err = gorm.Open(sqlite.New(sqlite.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		panic("failed to create gorm binding")
	}
}

func GetDatabase() *gorm.DB {
	return db
}

func MigrateTables() error {
	err := db.AutoMigrate(&models.User{}, &models.Vehicle{}, &models.Refuel{})
	if err != nil {
		return err
	}
	return nil
}

func DatabaseHealthCheck() error {
	var value int
	result := db.Raw("SELECT 1+1").Scan(&value)
	if value != 2 || result.Error != nil {
		return errors.New("error querying database")
	}
	return nil
}
