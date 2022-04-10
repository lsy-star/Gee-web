package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	r := gee.New()

	r.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "this is a GET request, URL is:%q\n", r.URL)
	})
	r.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "this is a POST request, URL is:%q\n", r.URL)
	})

	r.Run(":9101")
}
