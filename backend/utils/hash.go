package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword takes a plaintext password and returns a bcrypt hash.
// bcrypt.DefaultCost provides a reasonable balance between security and performance.
// Returns the hashed password as a string and any error encountered.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost)
	// Debugging: prints hashed password in bytes
	//
	// THIS NEEDS TO BE REMOVED
	//
	fmt.Println("Hashed Password Testing: ", bytes)
	return string(bytes), err
}

// CheckPasswordHash compares a plaintext password with its bcrypt hash.
// Returns true if the password matches the hash, false otherwise.
func CheckPasswordHash(password, hash string) bool {
	fmt.Println("check password: ", password)
	fmt.Println("check hash: ", hash)

	//HashPassword(password)

	// Compare the hashed password with the plaintext password
	err := bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	)

	// Debugging: prints error if passwords do not match
	fmt.Println("err: ", err)
	return err == nil
}
