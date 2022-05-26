package chiapi

import (
	"log"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
)

func TestChi(t *testing.T) {
	go func() {
		NewChiServer()
	}()

	time.Sleep(time.Second)
	c := resty.New()
	resp, err := c.R().Get("http://localhost:3333/")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(resp.String())
}
