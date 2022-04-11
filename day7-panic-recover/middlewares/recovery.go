package middlewares

import (
	"gee"
	"log"
	"net/http"
)

func Recovery() gee.HandlerFunc {
	return func(ctx *gee.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("%s\n\n", err)
				ctx.Fail(http.StatusInternalServerError, "500 InternalServerError")
			}
		}()
		ctx.Next()
	}
}
