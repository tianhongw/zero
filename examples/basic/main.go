package main

import (
	"net/http"

	"github.com/tianhongw/zero"
)

func main() {
	app := zero.New()
	app.Use(zero.Logger())
	app.Use(zero.Recovery())

	app.GET("/", func(c *zero.Context) {
		c.String(http.StatusOK, "Hello World!\n")
	})

	app.GET("/hello/:name", func(c *zero.Context) {
		c.String(http.StatusOK, "Hello %s!\n", c.Param("name"))
	})

	app.GET("/query", func(c *zero.Context) {
		c.JSON(http.StatusOK, zero.H{
			"name": "tianhongw",
			"age":  "25",
		})
	})

	v1 := app.Group("v1")
	{
		v1.GET("/", func(c *zero.Context) {
			c.String(http.StatusOK, c.Request.URL.Path)
		})

		v1.GET("/hello", func(c *zero.Context) {
			c.String(http.StatusOK, c.Request.URL.Path)
		})
		v1.GET("/panic", func(c *zero.Context) {
			a := make([]int, 1)
			a[10] = 1
			c.String(http.StatusOK, c.Request.URL.Path)
		})
	}

	app.Run(":8080")
}
