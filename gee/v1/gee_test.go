package gee

import (
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
)

func TestV1(t *testing.T) {
	go func() {
		NewV1Server()
	}()

	time.Sleep(time.Second)
	c := resty.New()
	resp, err := c.R().Get("http://localhost:5555/")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(resp.String())

	resp, err = c.R().Get("http://localhost:5555/hello")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(resp.String())
}

func TestV2(t *testing.T) {
	go func() {
		NewServerV2()
	}()

	time.Sleep(time.Second)
	c := resty.New()
	resp, err := c.R().Get("http://localhost:5555/")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(resp.String())

	resp, err = c.R().Get("http://localhost:5555/hello")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(resp.String())

	resp, err = c.R().Get("http://localhost:5555/404_not_found")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(resp.String())
}

func TestGee(t *testing.T) {
	go func() {
		r := New()
		r.Get("/", func(w http.ResponseWriter, req *http.Request) {
			w.Write([]byte("gee,get, hello"))
		})

		r.Post("/", func(w http.ResponseWriter, req *http.Request) {
			w.Write([]byte("gee,,post, hello"))
		})
		r.Run(":5555")
	}()

	time.Sleep(time.Second)
	c := resty.New()
	resp, err := c.R().Get("http://localhost:5555/")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(resp.String())

	resp, err = c.R().Post("http://localhost:5555/")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(resp.String())

	resp, err = c.R().Get("http://localhost:5555/404_not_found")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(resp.String())
}
