package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func author(u, p string) (string, string) {
	if u == "admin" && p == "astimone" {
		fmt.Println("Authorized!")
		return u, p
	}
	fmt.Println("Credentials wrong!")
	return u, p
}

func startPage(c *gin.Context) {
	var person Person
	if c.ShouldBindQuery(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(200, "Your name is %s and your address is %s\n", person.Name, person.Address)
	c.String(200, "Success")
}

func main() {
	r := gin.Default()
	// var username string
	// var pwd string
	// fmt.Printf("Please input your username: ")
	// fmt.Scan(&username)
	// fmt.Printf("Please input your password: ")
	// fmt.Scan(&pwd)
	// author(username, pwd)
	r.Any("/testing", startPage)
	r.Run(":8080")
}
