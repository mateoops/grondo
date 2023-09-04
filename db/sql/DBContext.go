package sql

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var SqlDb *sql.DB

func InitDB() error {
	var err error
	SqlDb, err = sql.Open("sqlite3", "internal.db")
	if err != nil {
		panic(err)
	}

	return nil
}

func CloseDB() {
	if SqlDb != nil {
		SqlDb.Close()
	}
}

func TestConnection() {
	var sqliteVersion string
	err := SqlDb.QueryRow("SELECT sqlite_version()").Scan(&sqliteVersion)
	if err != nil {
		panic(err)
	}

	fmt.Printf("SQLite version: %s\n", sqliteVersion)

}

func CreateCronJobTable() {
	// Create a table if it doesn't exist
	var err error
	_, err = SqlDb.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            username TEXT,
            email TEXT
        )
    `)
	if err != nil {
		panic(err)
	}

}
