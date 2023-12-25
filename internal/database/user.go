package database

import (
	"database/sql"
	"fmt"
)

type User struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}

type UserNotFound struct{}

func (*UserNotFound) Error() string {
	return "user not found"
}

func CreateUser(name, password string) error {
	_, err := Db.Exec("INSERT INTO users (name, password) VALUES (?, ?)", name, password)
	if err != nil {
		return fmt.Errorf("couldn't insert user, %w", err)
	}

	return nil
}

func FindUserByName(name string) (User, error) {
	//TODO: refactor Scan :justatest:
	var user User
	err := Db.QueryRow("SELECT * FROM users WHERE name = ?", name).Scan(&user.Id, &user.Name, &user.Password, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, &UserNotFound{}
		}

		return user, err
	}

	return user, nil
}

func FindUserById(id int) (User, error) {
	//TODO: refactor Scan :justatest:
	var user User
	err := Db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.Id, &user.Name, &user.Password, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, &UserNotFound{}
		}

		return user, err
	}

	return user, nil
}
