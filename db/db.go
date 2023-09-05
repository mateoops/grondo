package sql

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("internal.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
