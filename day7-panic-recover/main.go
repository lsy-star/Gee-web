package main

import (
	"gee"
	"middlewares"
	"net/http"
)

func main() {
	r := gee.New()
	r.Use(middlewares.Logger(), middlewares.Recovery())

	r.GET("/", func(ctx *gee.Context) {
		ctx.String(http.StatusOK, "Hello,Gee!")
	})

	r.GET("/panic", func(ctx *gee.Context) {
		a := 0
		ctx.String(http.StatusOK, "%d", 1/a)
	})

	r.Run(":9700")
}
