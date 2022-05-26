package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/yiGmMk/learn-web-framework/beegov2/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
