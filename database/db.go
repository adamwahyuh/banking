package database

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("database/db.sqlite"), &gorm.Config{})
	if err != nil {
		log.Fatal("Cant Load Db")
	}

	DB = db
	fmt.Print("Database Ok!\n")
}
