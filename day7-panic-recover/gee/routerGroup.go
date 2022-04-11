package gee

type RouterGroup struct {
	prefix      string        //路由组匹配前缀
	middlewares []HandlerFunc //中间件
	parent      *RouterGroup  //支持嵌套的路由分组
	engine      *Engine       //所有的路由分组都共享一个Engine实例
}

func (rg *RouterGroup) Group(prefix string) *RouterGroup {
	engine := rg.engine
	newRouterGroup := &RouterGroup{
		prefix: rg.prefix + prefix,
		parent: rg,
		engine: engine,
		//不要加这个，加了的话每一个新的子分组就会继承原来的父分组的所有中间件。
		//middlewares: rg.middlewares,
	}
	engine.groups = append(engine.groups, newRouterGroup)

	return newRouterGroup
}

func (rg *RouterGroup) addRoute(method string, pattern string, handler HandlerFunc) {
	rg.engine.router.addRoute(method, rg.prefix+pattern, handler)
}

func (rg *RouterGroup) GET(pattern string, handler HandlerFunc) {
	rg.addRoute("GET", pattern, handler)
}

func (rg *RouterGroup) POST(pattern string, handler HandlerFunc) {
	rg.addRoute("POST", pattern, handler)
}

func (rg *RouterGroup) Use(middlewares ...HandlerFunc) {
	rg.middlewares = append(rg.middlewares, middlewares...)
}

func (rg *RouterGroup) GetMiddlewares() []HandlerFunc {
	return rg.middlewares
}
