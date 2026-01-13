package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost)
	fmt.Println("Hashed Password Testing: ", bytes)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	fmt.Println("check password: ", password)
	fmt.Println("check hash: ", hash)
	HashPassword(password)

	err := bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	)
	fmt.Println("err: ", err)
	return err == nil
}
