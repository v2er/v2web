package routers

import (
	"v2web/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{}, "get:Home")
	beego.Router("/t/:id", &controllers.IndexController{}, "get:Topic")
}
