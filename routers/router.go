package routers

import (
	"Golang_RPG/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/api/", &controllers.MainController{})
	beego.Router("/api/login", &controllers.LoginController{})
}
