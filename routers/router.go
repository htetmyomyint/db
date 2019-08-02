package routers

import (
	"hw/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/db", &controllers.DBController{}, "*:Delete")
}
