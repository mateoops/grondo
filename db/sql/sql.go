package sql

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("internal.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		log.Println("Database connection established")
	}
}
