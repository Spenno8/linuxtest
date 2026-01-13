package main

import (
	//"fmt"
	//"log"
	"backend/config"
	"backend/routes"
	//"golang.org/x/crypto/bcrypt"
)

func main() {
	r := routes.SetupRouter() // Returns *gin.Engine
	config.InitDB()           // Initialize DB connection
	r.Run(":8080")

	// password := "passworduser3"
	// hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(hash))

}
