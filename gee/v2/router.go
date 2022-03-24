package gee

import "net/http"

// 路由表,非常简单的map结构存储了路由表，使用map存储键值对，索引非常高效
// 但是有一个弊端，键值对的存储的方式，只能用来索引静态路由
type router struct {
	handlers map[string]HandleFunc
}

func (r *router) addRoute(method, pattern string, handler HandleFunc) {
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandleFunc)}
}
