package middlewares

import (
	"fmt"
	"gee"
	"log"
	"time"
)

func Logger() gee.HandlerFunc {
	return func(ctx *gee.Context) {
		fmt.Println("---------------------------------------------------------------------------------------------")
		startTime := time.Now()
		ctx.Next()
		log.Printf("[%d] %s in %v\n", ctx.StatusCode, ctx.Request.RequestURI, time.Since(startTime))
	}
}

func V1Logger() gee.HandlerFunc {
	return func(ctx *gee.Context) {
		now := time.Now()
		log.Printf("this is for V1 group.\n [%d] URL:%s  |  RequestURI:%s in %v for group V1\n", ctx.StatusCode, ctx.Request.URL, ctx.Request.RequestURI, time.Since(now))
		ctx.Next()
	}
}

func SayHi() gee.HandlerFunc {
	return func(ctx *gee.Context) {
		log.Println("say Hi~~~")
		ctx.Next()
	}
}
