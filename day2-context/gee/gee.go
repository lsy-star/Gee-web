package gee

import (
	"net/http"
)

//http.HandlerFunc
type HandlerFunc func(ctx *Context)

type Engine struct {
	router *Router
}

func New() *Engine {
	return &Engine{router: NewRouter()}
}

func (e *Engine) addRouter(method string, pattern string, handler HandlerFunc) {
	e.router.addRouter(method, pattern, handler)
}

func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRouter("GET", pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRouter("POST", pattern, handler)
}

func (e *Engine) Run(address string) error {
	return http.ListenAndServe(address, e)
}

func (e *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	c := newContext(writer, request)
	e.router.handle(c)
}
