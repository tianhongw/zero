package main

import (
	"net/http"

	"github.com/trytwice/zero"
)

func main() {
	app := zero.New()

	app.GET("/", func(c *zero.Context) {
		c.String(http.StatusOK, "Hello World!\n")
	})

	app.GET("/hello", func(c *zero.Context) {
		c.String(http.StatusOK, "Hello %s!\n", c.Query("name"))
	})

	app.GET("/query", func(c *zero.Context) {
		c.JSON(http.StatusOK, zero.H{
			"name": "tianhongw",
			"age":  "25",
		})
	})

	app.Run(":8080")
}
