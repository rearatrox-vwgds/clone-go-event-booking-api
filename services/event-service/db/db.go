package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "event-api.db")

	if err != nil {
		panic("could not connect to the database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		creator_id INTEGER,
		FOREIGN KEY(creator_id) REFERENCES users(id)
	)`
	_, err := DB.Exec(createEventsTable)

	if err != nil {
		fmt.Println(err)
		panic("could not create events table")
	}

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS event_registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id)
	)`

	_, err = DB.Exec(createRegistrationsTable)
	if err != nil {
		fmt.Println(err)
		panic("could not create registrations table")
	}
}
