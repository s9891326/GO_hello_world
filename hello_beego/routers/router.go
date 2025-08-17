package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"hello_beego/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hello", &controllers.MainController{}, "get:Hello")
}
