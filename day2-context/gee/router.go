package gee

import (
	"log"
	"net/http"
)

type Router struct {
	handlers map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{handlers: map[string]HandlerFunc{}}
}

func (r *Router) addRouter(method string, pattern string, handler HandlerFunc) {
	log.Printf("Router : %4s - %s\n", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *Router) handle(ctx *Context) {
	key := ctx.Method + "-" + ctx.Path
	if handler, ok := r.handlers[key]; ok {
		handler(ctx)
	} else {
		ctx.String(http.StatusNotFound, "404 NOT FOUND : %s", ctx.Path)
	}

}
