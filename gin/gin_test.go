package ginapi

import (
	"log"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
)

func TestChi(t *testing.T) {
	go func() {
		NewGinServer()
	}()

	time.Sleep(time.Second)
	c := resty.New()
	resp, err := c.R().Post("http://localhost:4444/v2/gin/")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(resp.String())
}
