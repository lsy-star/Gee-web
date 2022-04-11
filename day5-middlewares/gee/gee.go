package gee

import (
	"fmt"
	"net/http"
	"strings"
)

type HandlerFunc func(ctx *Context)

type Engine struct {
	*RouterGroup //这里我不确定，是表示Engine本身也作为一个顶层的路由分组吗？
	router       *router
	groups       []*RouterGroup //存储所有的路由分组
}

func New() *Engine {
	e := &Engine{router: newRouter()}
	e.RouterGroup = &RouterGroup{engine: e}
	e.groups = []*RouterGroup{e.RouterGroup} //这一步是什么意思?
	//是向slice里面赋一个元素，类似于[]int{1,2,3}这样赋初值
	return e
}

func (e *Engine) Run(address string) error {
	return http.ListenAndServe(address, e)
}

func (e *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range e.groups {
		//加一个"/"是为了避免请求路径："v11"和当前分组路径前缀："/v1"匹配上
		if strings.HasPrefix(request.URL.Path, group.prefix+"/") {
			fmt.Println(request.URL.Path)
			fmt.Println(group.prefix)
			fmt.Println("++++++++++++")
			//if request.URL.Path == group.prefix+"/" {
			middlewares = append(middlewares, group.middlewares...)
		}
	}

	c := newContext(writer, request)
	c.handlers = middlewares
	e.router.handle(c)
}

func (e *Engine) GetMiddlewares() []HandlerFunc {
	return e.middlewares
}
