package main

import (
	"database/sql"
	"log"
	"time"

	"forum/entity"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Open SQLite3 database
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create User table if not exists
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT UNIQUE,
			password TEXT,
			email TEXT UNIQUE,
			created_at TIMESTAMP,
			updated_at TIMESTAMP
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Example usage:
	// Create a new user
	newUser := entity.User{
		Username:  "john_doe",
		Password:  "password123",
		Email:     "john@example.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
