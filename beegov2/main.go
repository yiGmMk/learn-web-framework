package main

import (
	_ "bee-ocr/routers"

	"bee-ocr/util"

	"github.com/beego/beego/v2/core/admin"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	// admin ,add healthcheck
	admin.AddHealthCheck("health-cpu", &util.CpuCheck{})

	beego.Run()
}
