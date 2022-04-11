## day5-middlewares（2022.4.11）

---

### 内容

- 设计并实现Web框架的中间件机制
- 实现一个超简单的通用`Logger`中间件(也就三两行代码)

### 步骤

- 每一个路由分组`RouterGroup`对象都有一个属性`middlewares`（类型为`[]HandlerFunc`）存储该分组的所有的中间件
- 进行路由匹配的时候，请求路径每匹配到一个路由分组就会把该路由分组的所有中间件加载到一个`slice`中，之后将该`slice`作为一个`Context`类型对象的`handlers`属性值，再把该`Context`
  对象交给全局的`Engine`实例（也就是`main`函数中创建的`r := gee.New()`）执行

### bug

在一个地方写错了，找了几个小时的bug

```go
v1 := r.Group("/v1")
 ```

这里我写成`"v1"`了，furious
