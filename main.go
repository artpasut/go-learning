package main

import "github.com/gin-gonic/gin"

func homepage(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"title": "Home Page",
	})
}

func main() {
	app := gin.Default()
	app.LoadHTMLGlob("templates/*")
	app.GET("/", homepage)
	app.Run(":8080")
}
