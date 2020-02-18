package main

import (
	"fmt"
	"net/http"

	"github.com/trytwice/zero"
)

func main() {
	zero := zero.New()

	zero.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Hello World!\n")
	})

	zero.GET("/ping", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "pong\n")
	})

	zero.Run(":8080")
}
