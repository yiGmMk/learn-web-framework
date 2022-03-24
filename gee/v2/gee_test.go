package gee

import (
	"log"
	"testing"

	"github.com/go-resty/resty/v2"
)

func TestGee(t *testing.T) {
	go func() {
		r := New()
		r.Get("/", func(c *Context) {
			c.JSON(200, H{
				"hello":  "world",
				"method": c.Req.Method,
				"name":   c.Req.URL.Query().Get("name")})
		})

		r.Post("/", func(c *Context) {
			c.JSON(200, H{
				"hello":  "world",
				"method": c.Req.Method,
				"name":   c.Req.URL.Query().Get("name")})
		})
		r.Run()
	}()

	c := resty.New().R()
	resp, err := c.Post("http://localhost:9999/?name=123456")
	if err != nil {
		log.Println(err)
	} else {
		log.Println(resp.String())
	}
	resp, err = c.Get("http://localhost:9999/?name=yiGmMk")
	if err != nil {
		log.Println(err)
	} else {
		log.Println(resp.String())
	}
}
