package gozero

import (
	"fmt"

	"github.com/zeromicro/go-zero/rest"
)

func Test() {
	c := rest.RestConf{}
	fmt.Println(c)

	server := rest.MustNewServer(c)
	server.Start()
}
