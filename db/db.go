package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Database variable
var DB *sql.DB

func InitDB() {
	//Open our sql
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Could not connect to database.")
	}
	err = DB.Ping()
	if err != nil {
		panic("Ping failed")
	}
	//Set max open/idle connections
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
}

func createTables() {
	createEventsTable :=
		`CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER
	)`
	fmt.Println(createEventsTable)
	//Create the events table
	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic("Could not create events table.")
	}
}