package gee

import (
	"log"
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

//roots key eg, roots['GET'] roots['POST']
//handlers key eg, handlers['GET-/p/:lang/doc'] handlers['POST-/p/book']

func newRouter() *router {
	return &router{
		roots:    map[string]*node{},
		handlers: map[string]HandlerFunc{},
	}
}

func parsePattern2Parts(pattern string) (res []string) {
	split := strings.Split(pattern, "/")

	for _, part := range split {
		if part != "" {
			res = append(res, part)
			if part[0] == '*' {
				break
			}
		}
	}
	return
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Router : %4s - %s\n", method, pattern)

	key := method + "-" + pattern

	if _, ok := r.roots[method]; !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parsePattern2Parts(pattern), 0)
	r.handlers[key] = handler
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	params := map[string]string{}
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	searchedParts := parsePattern2Parts(path)
	n := root.search(searchedParts, 0)
	if n != nil {
		parts := parsePattern2Parts(n.pattern)
		for i, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchedParts[i]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchedParts[i:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

func (r *router) handle(ctx *Context) {
	n, params := r.getRoute(ctx.Method, ctx.Path)
	if n != nil {
		ctx.Params = params
		key := ctx.Method + "-" + n.pattern
		ctx.handlers = append(ctx.handlers, r.handlers[key])
	} else {
		ctx.handlers = append(ctx.handlers, func(ctx *Context) {
			ctx.String(http.StatusNotFound, "404 NOT FOUND : %s ", ctx.Path)
		})
	}
	ctx.Next()
}
