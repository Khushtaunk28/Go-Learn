package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	 var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns((5))

	createTables()
}


func createTables() {
	if DB == nil {
		panic("Database connection is not initialized")
	}

	createUserTables:=`
	CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL
	)`
	_,err:= DB.Exec(createUserTables)
	if err != nil {
		panic(fmt.Sprintf("Could not create users table: %v", err)) // Print actual error
	}



	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		if err != nil {
			panic(fmt.Sprintf("Could not create events table: %v", err)) // Print actual error
		}
	}


	createRegistrationTable:=`
	CREATE TABLE IF NOT EXISTS registration(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	event_id INTEGER,
	user_id INTEGER,
	FOREIGN KEY(event_id) REFERENCES events(id),
	FOREIGN KEY(user_id) REFERENCES users(id),
	)
	`
	_,err=DB.Exec(createRegistrationTable)
	if err != nil {
		if err != nil {
			panic(fmt.Sprintf("Could not create registration table: %v", err)) // Print actual error
		}
	}

}