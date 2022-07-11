package chiapi

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-resty/resty/v2"
)

type Human struct {
	Name string
}

var _people = Human{Name: "zhangsan"}

func TestFmt(t *testing.T) {
	s := fmt.Sprintf("%q", "Hello World")
	fmt.Println(s)
	fmt.Printf("%v\n%+v\n%#v", _people, _people, _people)

	err := errors.New("error")
	e1, e2 := fmt.Errorf("err:%w", err), fmt.Errorf("err:%v", err)
	fmt.Println(e1, e2)
}

func TestChi(t *testing.T) {
	go func() {
		m := chi.NewRouter()
		m.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(fmt.Sprintf("%+v", _people)))
		})
		http.ListenAndServe(":3333", m)
	}()

	time.Sleep(time.Second)
	c := resty.New()
	resp, err := c.R().Get("http://localhost:3333/")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("chi resp", resp.String())
}
