package model

import (
	"backend/config"
	"context"
	"fmt"
)

// User represents a user in the system.
// Password is stored as a hashed string, not plaintext.
type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string
}

// GetUserByCred retrieves a user from the database using a specific credential.
// credtype must be either "email" or "username".
// Returns a pointer to a User or an error if not found.
func GetUserByCred(credtype string, cred string) (*User, error) {
	fmt.Println("QUERYING USER:", credtype, " FOR: ", cred)

	var u User
	var column string

	// Determine which column to search by
	switch credtype {
	case "email":
		column = "email"
	case "username":
		column = "username"
	default:
		return nil, fmt.Errorf("invalid credential type")
	}

	// Construct the SQL query dynamically
	query := fmt.Sprintf(
		"SELECT id, email, username, password_hash FROM users WHERE %s = $1",
		column,
	)
	// Execute query
	row := config.DB.QueryRow(context.Background(), query, cred)

	// Scan results into User struct
	err := row.Scan(&u.ID, &u.Email, &u.Username, &u.Password)
	if err != nil {
		fmt.Println("DB ERROR:", err)
		return nil, err
	}

	fmt.Println("Users Hash from DB:", u.Password)

	return &u, nil
}

// SignupNewUser creates a new user in the database.
// Accepts email, username, hashed password, firstname, and lastname.
// Returns the newly created User with its ID populated.
func SignupNewUser(email, username, password, firstname, lastname string) (*User, error) {
	var u User

	// Insert new user into the database
	row := config.DB.QueryRow(context.Background(),
		"INSERT INTO users (username, email, first_name, last_name, password_hash) VALUES ($1, $2, $3, $4, $5) RETURNING id, email, username",
		username, email, firstname, lastname, password)

	// Scan returned values into User struct
	err := row.Scan(&u.ID, &u.Email, &u.Password)
	if err != nil {
		fmt.Println("DB ERROR:", err)
		return nil, err
	}
	return &u, nil
}
