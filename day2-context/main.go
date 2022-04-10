package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()

	r.GET("/hello", func(ctx *gee.Context) {
		//ctx.HTML(http.StatusOK, "<h1>Hello</h1>")
		ctx.String(http.StatusOK, "hello %s, you are at %s\n", ctx.QueryString("name"), ctx.Path)
	})

	r.GET("/", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "<h1>Gee!</h1>")
	})

	r.POST("/login", func(ctx *gee.Context) {
		ctx.JSON(http.StatusOK, gee.H{
			"username": ctx.Form("username"),
			"password": ctx.Form("password"),
		})
	})

	r.Run(":9200")

}
