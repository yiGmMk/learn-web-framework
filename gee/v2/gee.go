package gee

import (
	"fmt"
	"net/http"
)

// references: https://geektutu.com/post/gee-day2.html
/*
实现下面的效果
func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *gee.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
Handler的参数变成成了gee.Context，提供了查询Query/PostForm参数的功能。
gee.Context封装了HTML/String/JSON函数，能够快速构造HTTP响应。

对Web服务来说，无非是根据请求*http.Request，构造响应http.ResponseWriter。但是这两个对象提供的接口粒度太细，
比如我们要构造一个完整的响应，需要考虑消息头(Header)和消息体(Body)，而 Header 包含了状态码(StatusCode)，消息类型(ContentType)等
几乎每次请求都需要设置的信息。因此，如果不进行有效的封装，那么框架的用户将需要写大量重复，繁杂的代码，而且容易出错。
针对常用场景，能够高效地构造出 HTTP 响应是一个好的框架必须考虑的点。
用返回 JSON 数据作比较，感受下封装前后的差距。

封装前
obj = map[string]interface{}{
    "name": "geektutu",
    "password": "1234",
}
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusOK)
encoder := json.NewEncoder(w)
if err := encoder.Encode(obj); err != nil {
    http.Error(w, err.Error(), 500)
}
VS 封装后：
c.JSON(http.StatusOK, gee.H{
    "username": c.PostForm("username"),
    "password": c.PostForm("password"),
})
*/

type HandleFunc func(http.ResponseWriter, *http.Request)

// 2.实现http.Handler接口,拦截所有的HTTP请求，拥有了统一的控制入口。在这里我们可以自由定义路由映射的规则，也可以统一添加一些处理逻辑，例如日志、异常处理等。
type Engine struct {
	router map[string]HandleFunc
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
