package gee

import (
	"log"
	"net/http"
	"strings"
)

// 通过前缀树实现动态路由=>实现路由匹配和一些path参数的解析
type router struct {
	roots    map[string]*node
	handlers map[string]HandleFunc
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandleFunc),
		roots:    make(map[string]*node),
	}
}

// url解析, 只允许一个*,
func parsePattern(pattern string) []string {
	ps := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, item := range ps {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

// 新增路由,构造前缀树分支
func (r *router) addRoute(method, pattern string, handler HandleFunc) {
	ps := parsePattern(pattern)

	if _, ok := r.roots[method]; !ok {
		r.roots[method] = &node{}
	}

	r.roots[method].insert(pattern, ps, 0)

	key := method + "-" + pattern
	r.handlers[key] = handler
}

// 路由匹配,和path参数处理
// /user/:id/profile
func (r *router) getRoute(method, path string) (*node, map[string]string) {
	parts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	// 匹配到的路径叶子节点
	node := root.search(parts, 0)
	if node == nil {
		return nil, nil
	}
	nodeParts := parsePattern(node.pattern)
	for i, part := range nodeParts {
		if part[0] == ':' {
			params[part[1:]] = parts[i]
		}
		if part[0] == '*' && len(part) > 1 {
			params[part[1:]] = strings.Join(parts[i:], "/")
			break
		}
	}
	return node, params
}

// 请求处理函数,在ServeHTTP里中调用
func (r *router) handle(c *Context) {
	node, params := r.getRoute(c.Method, c.Path)
	if node == nil {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		return
	}
	// 动态路由这里HandleFunc映射的是的pattern不能直接使用c.Path,c.Path是请求参数,有一些是参数值
	/* 如 /hello/:name   => /hello/gee  	 */
	log.Printf("node.pat:%s,c.pat:%s\n", node.pattern, c.Path)
	key := c.Method + "-" + node.pattern
	c.Params = params
	r.handlers[key](c)
}
