package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	fmt.Println(string(password))
}
