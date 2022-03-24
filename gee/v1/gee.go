package gee

import (
	"fmt"
	"net/http"
)

// 1.通过HandleFunc添加路由,调用 http.HandleFunc 实现了路由和Handler的映射
func NewV1Server() error {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	return http.ListenAndServe(":5555", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "url path=%q\n", r.URL.Path)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "post,hello,below is the header:")
	for k, v := range r.Header {
		fmt.Fprintf(w, "%s:%s\n", k, v)
	}
}

type HandleFunc func(http.ResponseWriter, *http.Request)

// 2.实现http.Handler接口,拦截所有的HTTP请求，拥有了统一的控制入口。在这里我们可以自由定义路由映射的规则，也可以统一添加一些处理逻辑，例如日志、异常处理等。
type Engine struct {
	router map[string]HandleFunc
}

// func (e Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		fmt.Fprintf(w, "url path=%q\n", r.URL.Path)
// 	case "/hello":
// 		fmt.Fprintln(w, "post,hello,below is the header:")
// 		for k, v := range r.Header {
// 			fmt.Fprintf(w, "%s:%s\n", k, v)
// 		}
// 	default:
// 		fmt.Fprintf(w, "404 not found,url:%s", r.URL.Path)
// 	}
// }

func NewServerV2() error {
	return http.ListenAndServe(":5555", Engine{})
}

// 3.构造Engine
func New() *Engine {
	return &Engine{router: make(map[string]HandleFunc)}
}

func (e *Engine) addRouter(method, pattern string, handler HandleFunc) {
	key := method + "-" + pattern
	e.router[key] = handler
}

func (e *Engine) Get(pattern string, handler HandleFunc) {
	e.addRouter("GET", pattern, handler)
}

func (e *Engine) Post(pattern string, handler HandleFunc) {
	e.addRouter("POST", pattern, handler)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

// 解析请求路径,查找路由映射表,执行对应的处理函数
func (e Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := e.router[key]; ok {
		handler(w, r)
	} else {
		fmt.Fprintf(w, "404 not found,url:%s", r.URL.Path)
	}
}
