package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Connect to data base failed")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTable()
}

func createTable() {

	createUserTableQry := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE ,
			password TEXT NOT NULL
		)
	`

	_, err := DB.Exec(createUserTableQry)
	if err != nil {
		log.Fatalf("User table create fail: %v", err)
	}

	createEventTableQry := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			dateTime DATETIME NOT NULL,
			userId INTEGER,
			FOREIGN KEY (userId) REFERENCES users(id)
		)
	`
	_, err = DB.Exec(createEventTableQry)

	if err != nil {
		log.Fatalf("Event table create fail: %v", err)
	}
}
