package main

import (
	"golang.org/x/crypto/bcrypt"
	"fmt"
	"os"
)

func main() {
	res, err := bcrypt.GenerateFromPassword([]byte(os.Args[1]), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))
}
