package repo

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type IDb interface {
	IUser
	IPost
	IComment
	IProfile
}
type repo struct {
	db  *sql.DB
	log *log.Logger
}

func New() (IDb, error) {
	// Open SQLite3 database
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Printf("error while create table of database %s\n", err.Error())
		return nil, err
	}
	err = createTables(db)
	if err != nil {
		log.Printf("error while create table of database %s\n", err.Error())
		return nil, err
	}
	l := &log.Logger{}

	return repo{db: db, log: l}, nil
}

// createTables creates User, Post, Comment, and Profile tables
func createTables(db *sql.DB) error {
	fmt.Println("create db")
	// Create User table
	_, err := db.Exec(`
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
		return err
	}

	// Create Post table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS posts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER,
			title TEXT,
			content TEXT,
			created_at TIMESTAMP,
			updated_at TIMESTAMP,
			likes INTEGER DEFAULT 0,
			dislikes INTEGER DEFAULT 0,
			FOREIGN KEY (user_id) REFERENCES users(id)
		)
	`)
	if err != nil {
		return err
	}

	// Create Comment table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS comments (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER,
			post_id INTEGER,
			content TEXT,
			created_at TIMESTAMP,
			updated_at TIMESTAMP,
			likes INTEGER DEFAULT 0,
			dislikes INTEGER DEFAULT 0,
			FOREIGN KEY (user_id) REFERENCES users(id),
			FOREIGN KEY (post_id) REFERENCES posts(id)
		)
	`)
	if err != nil {
		return err
	}

	// Create Profile table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS profiles (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER,
			name TEXT,
			bio TEXT,
			image_url TEXT,
			created_at TIMESTAMP,
			updated_at TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id)
		)
	`)
	if err != nil {
		return err
	}

	return nil
}
