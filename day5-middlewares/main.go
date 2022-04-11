package main

import (
	"gee"
	"middlewares"
	"net/http"
)

func main() {
	r := gee.New()
	r.Use(middlewares.Logger(), middlewares.SayHi())

	r.GET("/", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "<h1>hello Gee!</h1>")
	})

	v1 := r.Group("/v1") //这里我写成"v1"了，找了我几个小时的bug。furious
	v1.Use(middlewares.V1Logger())

	{
		v1.GET("/hello/:name", func(ctx *gee.Context) {
			ctx.String(http.StatusOK, "dynamic route, Hello %s, you are at %s\n", ctx.Param("name"), ctx.Path)
		})
	}

	r.Run(":9500")

}

//func onlyForV2() gee.HandlerFunc {
//	return func(c *gee.Context) {
//		// Start timer
//		t := time.Now()
//		// if a server error occurred
//		c.Fail(500, "Internal Server Error")
//		// Calculate resolution time
//		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Request.RequestURI, time.Since(t))
//	}
//}
//
//func main() {
//	r := gee.New()
//	r.Use(gee.Logger()) // global midlleware
//	r.GET("/", func(c *gee.Context) {
//		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
//	})
//
//	v2 := r.Group("/v2")
//	//v2.Use(onlyForV2()) // v2 group middleware
//	v2.Use(func(ctx *gee.Context) {
//		fmt.Println("Say hi~~~~~~~~~~")
//	}, func(ctx *gee.Context) {
//		fmt.Println("12345678790")
//	})
//	{
//		v2.GET("/hello/:name", func(c *gee.Context) {
//			// expect /hello/geektutu
//			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
//		})
//	}
//
//	fmt.Println(r.GetMiddlewares())
//	fmt.Println(v2.GetMiddlewares())
//
//	r.Run(":9999")
//}
