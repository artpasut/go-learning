package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func main() {
	c := add(5, 6)
	var id string
	var password string
	fmt.Println("Hello, This is my first time with go.")
	fmt.Println("Function add return value", c)
	fmt.Print("Enter USERNAME: ")
	fmt.Scanf("%s", &id)
	fmt.Print("Enter PASSWORD: ")
	fmt.Scanf("%s", &password)
	fmt.Println("Your USERNAME is", id)
	if id == "admin" && password == "admin@123" {
		fmt.Println("Login Successfully!")
	} else {
		fmt.Println("Credentials wrong")
	}
}
