package gee

import (
	"net/http"
)

type HandlerFunc func(ctx *Context)

type (
	RouterGroup struct {
		prefix      string        //路由组匹配前缀
		middlewares []HandlerFunc //中间件
		parent      *RouterGroup  //支持嵌套的路由分组
		engine      *Engine       //所有的路由分组都共享一个Engine实例
	}

	Engine struct {
		*RouterGroup //这里我不确定，是表示Engine本身也作为一个顶层的路由分组吗？
		router       *router
		groups       []*RouterGroup //存储所有的路由分组
	}
)

func New() *Engine {
	e := &Engine{router: newRouter()}
	e.RouterGroup = &RouterGroup{engine: e}
	e.groups = []*RouterGroup{e.RouterGroup} //这一步是什么意思?
	//是向slice里面赋一个元素，类似于[]int{1,2,3}这样赋初值
	return e
}

func (rg *RouterGroup) Group(prefix string) *RouterGroup {
	engine := rg.engine
	newRouterGroup := &RouterGroup{
		prefix: rg.prefix + prefix,
		parent: rg,
		engine: engine,
	}
	engine.groups = append(engine.groups, newRouterGroup)
	return newRouterGroup
}

func (rg *RouterGroup) addRouter(method string, pattern string, handler HandlerFunc) {
	rg.engine.router.addRoute(method, rg.prefix+pattern, handler)
}

func (rg *RouterGroup) GET(pattern string, handler HandlerFunc) {
	rg.addRouter("GET", pattern, handler)
}

func (rg *RouterGroup) POST(pattern string, handler HandlerFunc) {
	rg.addRouter("POST", pattern, handler)
}

func (e *Engine) Run(address string) error {
	return http.ListenAndServe(address, e)
}

func (e *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	c := newContext(writer, request)
	e.router.handle(c)
}
