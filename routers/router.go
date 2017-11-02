package routers

import (
	"Golang_RPG/controllers"
	"Golang_RPG/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	// TODO: put initializations in a differnt file
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// TODO: add connection string in app.con and unify it on our systems
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("connectionString"))
	orm.RegisterModel(new(models.Users))
	orm.RegisterModel(new(models.Bots))
	beego.Router("/api/", &controllers.MainController{})
	beego.Router("/api/register", &controllers.RegisterController{})
	beego.Router("/api/login", &controllers.LoginController{})
	beego.Router("/api/bot", &controllers.BotController{})
}
