package main

import (
	"fmt"
)

func author(u, p string) (string, string) {
	if u == "admin" && p == "astimone" {
		fmt.Println("Authorized!")
		return u, p
	}
	fmt.Println("Credentials wrong!")
	return u, p
}

func main() {
	var username string
	var pwd string
	fmt.Printf("Please input your user name: ")
	fmt.Scan(&username)
	fmt.Printf("Please input your password: ")
	fmt.Scan(&pwd)
	author(username, pwd)
}
