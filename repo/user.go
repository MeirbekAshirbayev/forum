package repo

import (
	"forum/entity"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type IUser interface {
	CreateUser(user *entity.User) error
	GetUserByID(id int64) (*entity.User, error)
	UpdateUser(user entity.User) error
	DeleteUserByID(id int64) error
}

// createUser creates a new user in the User table
func (r repo) CreateUser(user *entity.User) error {
	log.Println("k")

	stmt, err := r.db.Prepare("INSERT INTO users (username, password, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		r.log.Printf("error while to prepare datas to write into the user table: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.Password, user.Email, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		r.log.Printf("error while exec prepared datas to write into user table: %s\n", err.Error())
		return err
	}

	return nil
}

// getUserByID retrieves a user from the User table by ID
func (r repo) GetUserByID(id int64) (*entity.User, error) {
	stmt, err := r.db.Prepare("SELECT id, username, password, email, created_at, updated_at FROM users WHERE id = ?")
	if err != nil {
		r.log.Printf("error while to prepare datas to get user by id from user table: %s\n", err.Error())
		return nil, err
	}
	defer stmt.Close()

	var user entity.User
	err = stmt.QueryRow(id).Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		r.log.Printf("error while to query row and scan user to get by id: %s\n", err.Error())
		return nil, err
	}

	return &user, nil
}

// updateUser updates an existing user in the User table
func (r repo) UpdateUser(user entity.User) error {
	stmt, err := r.db.Prepare("UPDATE users SET username = ?, password = ?, email = ?, updated_at = ? WHERE id = ?")
	if err != nil {
		r.log.Printf("error while to prepare update datas in user table: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.Password, user.Email, user.UpdatedAt, user.Id)
	if err != nil {
		r.log.Printf("error while exec prepared update datas in user table: %s\n", err.Error())
		return err
	}

	return nil
}

// deleteUserByID deletes a user from the User table by ID
func (r repo) DeleteUserByID(id int64) error {
	stmt, err := r.db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		r.log.Printf("error while to prepare delete user by id in user table: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		r.log.Printf("error while exec prepared delete user by id in user table: %s\n", err.Error())
		return err
	}

	return nil
}
