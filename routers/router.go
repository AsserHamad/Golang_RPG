package routers

import (
	"Golang_RPG/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// TODO: put initializations in a differnt file

	beego.Router("/api/", &controllers.MainController{})
	beego.Router("/api/register", &controllers.RegisterController{})
	beego.Router("/api/login", &controllers.LoginController{})
	beego.Router("/api/bot", &controllers.BotController{})
}
