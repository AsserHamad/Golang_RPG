package routers

import (
	"Golang_RPG/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/", &controllers.MainController{})
	beego.Router("/api/register", &controllers.RegisterController{})
	beego.Router("/api/login", &controllers.LoginController{})
	beego.Router("/api/bot", &controllers.BotController{})
	beego.Router("/api/search", &controllers.ShopsSearchController{}) // searches for nearest shops
	beego.Router("/api/scan", &controllers.ScanController{})
}
