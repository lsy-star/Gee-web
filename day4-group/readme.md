## day4-group(2022.4.11)

---

###内容 
- 实现路由分组控制功能，支持路由分组嵌套
- 路由分组方便后续加入中间件（`day5-middlewares`完成）

###步骤
- 增加`RouterGroup`类，记录每一个路由分组
- `RouterGroup`同时嵌入`Engine`中，`Engine`作为整个框架顶层的分组，记录所有的路由分组信息。
- 在`RouterGroup`中加入`*Engine`属性，可以由任意一个路由分组到达作为顶层分组的`Engine`中。

###新增加结构体
```go
type (
	RouterGroup struct {
		prefix      string        //路由组匹配前缀
		middlewares []HandlerFunc //中间件
		parent      *RouterGroup  //支持嵌套的路由分组
		engine      *Engine       //所有的路由分组都共享一个Engine实例
	}

	Engine struct {
		*RouterGroup //这里我不确定，是表示Engine本身也作为一个顶层的路由分组吗？---应该是的
		router       *router
		groups       []*RouterGroup //存储所有的路由分组
	}
)
```
