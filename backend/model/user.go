package model

import (
	"backend/config"
	"context"
	"fmt"
)

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string
}

func GetUserByCred(credtype string, cred string) (*User, error) {
	fmt.Println("QUERYING USER:", credtype, " FOR: ", cred)

	var u User
	var column string
	switch credtype {
	case "email":
		column = "email"
	case "username":
		column = "username"
	default:
		return nil, fmt.Errorf("invalid credential type")
	}

	query := fmt.Sprintf(
		"SELECT id, email, username, password_hash FROM users WHERE %s = $1",
		column,
	)

	row := config.DB.QueryRow(context.Background(), query, cred)
	err := row.Scan(&u.ID, &u.Email, &u.Username, &u.Password)
	if err != nil {
		fmt.Println("DB ERROR:", err)
		return nil, err
	}

	fmt.Println("Users Hash from DB:", u.Password)

	return &u, nil
}

func SignupNewUser(email, username, password, firstname, lastname string) (*User, error) {
	var u User
	row := config.DB.QueryRow(context.Background(),
		"INSERT INTO users (username, email, first_name, last_name, password_hash) VALUES ($1, $2, $3, $4, $5) RETURNING id, email, username",
		username, email, firstname, lastname, password)
	err := row.Scan(&u.ID, &u.Email, &u.Password)
	if err != nil {
		fmt.Println("DB ERROR:", err)
		return nil, err
	}
	return &u, nil
}
