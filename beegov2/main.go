package main

import (
	"log"
	"time"

	_ "github.com/yiGmMk/learn-web-framework/beegov2/routers"

	"github.com/yiGmMk/learn-web-framework/beegov2/util"

	"github.com/beego/beego/v2/core/admin"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	// admin ,add healthcheck
	admin.AddHealthCheck("health-cpu", &util.CpuCheck{})
	log.Println("bee v2 start", time.Now().Format("2006-01-02 15:04:05"))
	beego.Run()
}
