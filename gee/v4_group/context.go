package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
针对使用场景，封装*http.Request和http.ResponseWriter的方法，简化相关接口的调用，只是设计 Context 的原因之一。
对于框架来说，还需要支撑额外的功能。例如，将来解析动态路由/hello/:name，参数:name的值放在哪呢？再比如，框架需要支持中间件，
那中间件产生的信息放在哪呢？Context 随着每一个请求的出现而产生，请求的结束而销毁，和当前请求强相关的信息都应由 Context 承载。
因此，设计 Context 结构，扩展性和复杂性留在了内部，而对外简化了接口。
路由的处理函数，以及将要实现的中间件，参数都统一使用 Context 实例， Context 就像一次会话的百宝箱，可以找到任何东西。
*/
type H map[string]interface{}

type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request
	// 请求
	Path   string
	Method string
	// 响应
	StatusCode int
	Params     map[string]string
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    r,
		Path:   r.URL.Path,
		Method: r.Method,
	}
}

// form表单参数
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

//  query参数
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// 路径参数
func (c *Context) Param(key string) string {
	return c.Params[key]
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// 设置响应头
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// 文本响应
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// 正确的调用顺序应该是Header().Set 然后WriteHeader() 最后是Write()
// http.ResponseWriter.Header().Set("Content-Type", "application/json")
// http.ResponseWriter.WriteHeader(http.StatusOK)
// http.ResponseWriter.Write([]byte(`{"message": "hello world"}`))
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
