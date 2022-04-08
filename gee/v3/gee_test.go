package gee

import (
	"log"
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"
)

/*
(1)
$ curl -i http://localhost:9999/
HTTP/1.1 200 OK
Date: Mon, 12 Aug 2019 16:52:52 GMT
Content-Length: 18
Content-Type: text/html; charset=utf-8
<h1>Hello Gee</h1>

(2)
$ curl "http://localhost:9999/hello?name=geektutu"
hello geektutu, you're at /hello

(3)
$ curl "http://localhost:9999/hello/geektutu"
hello geektutu, you're at /hello/geektutu

(4)
$ curl "http://localhost:9999/assets/css/geektutu.css"
{"filepath":"css/geektutu.css"}

(5)
$ curl "http://localhost:9999/xxx"
404 NOT FOUND: /xxx
*/

func TestGee(t *testing.T) {
	go func() {
		r := New()
		r.Get("/", func(c *Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		r.Get("/hello", func(c *Context) {
			// expect /hello?name=geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})

		r.Get("/hello/:name", func(c *Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})

		r.Get("/assets/*filepath", func(c *Context) {
			c.JSON(http.StatusOK, H{"filepath": c.Param("filepath")})
		})

		r.Run(":9999")
	}()

	c := resty.New().R()
	resp, err := c.Get("http://localhost:9999/")
	if err != nil {
		log.Println(err)
	} else {
		log.Println(resp.String())
	}
	resp, err = c.Get("http://localhost:9999/hello?name=yiGmMk")
	if err != nil {
		log.Println(err)
	} else {
		log.Println(resp.String())
	}
	resp, err = c.Get("http://localhost:9999/hello/yiGmMk")
	if err != nil {
		log.Println(err)
	} else {
		log.Println(resp.String())
	}
	resp, err = c.Get("http://localhost:9999/assets/name.txt")
	if err != nil {
		log.Println(err)
	} else {
		log.Println(resp.String())
	}
	resp, err = c.Get("http://localhost:9999/xxx")
	if err != nil {
		log.Println(err)
	} else {
		log.Println(resp.String())
	}
}
