package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()

	r.GET("/", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "<h1>any request!</h1>")
	})

	r.GET("/hello", func(ctx *gee.Context) {
		ctx.String(http.StatusOK, "hello %s, you are at %s\n", ctx.QueryString("name"), ctx.Path)
	})
	r.GET("/hello/:name", func(ctx *gee.Context) {
		ctx.String(http.StatusOK, "hello %s, dynamic patterning, you are at %s\n", ctx.Param("name"), ctx.Path)
	})

	r.GET("/assets/*filepath", func(ctx *gee.Context) {
		ctx.JSON(http.StatusOK, gee.H{
			"filepath": ctx.Param("filepath"),
		})
	})

	r.Run(":9300")
}
