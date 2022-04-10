package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//存放任意内容的map
type H map[string]interface{}

//封装和当前请求强相关的信息，扩展性和复杂性留在内部，对外简化接口
type Context struct {
	//原始的请求和响应
	Writer  http.ResponseWriter
	Request *http.Request

	//请求相关的信息
	Path   string
	Method string

	//响应相关的信息
	statusCode int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer:  w,
		Request: r,
		Path:    r.URL.Path,
		Method:  r.Method,
	}
}

func (c *Context) Form(key string) string {
	return c.Request.Form.Get(key)
}

func (c *Context) QueryString(key string) string {
	return c.Request.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.statusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) SetHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}
