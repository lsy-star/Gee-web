package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()

	r.GET("/index", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "<h1>Index Page!</h1>")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(ctx *gee.Context) {
			ctx.HTML(http.StatusOK, "<h2>Gee!</h2>")
		})

		v1.GET("/hello", func(ctx *gee.Context) {
			ctx.String(http.StatusOK, "hello %s, you are at %s\n", ctx.QueryString("name"), ctx.Path)
		})
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(ctx *gee.Context) {
			ctx.String(http.StatusOK, "(dynamic route) Hello %s, you are at %s\n", ctx.Params["name"], ctx.Path)
		})

		v2.POST("/login", func(ctx *gee.Context) {
			ctx.JSON(http.StatusOK, gee.H{
				"username": ctx.Form("username"),
				"password": ctx.Form("password"),
			})
		})
	}

	r.Run(":9400")
}
