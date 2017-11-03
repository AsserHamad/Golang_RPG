package routers

import (
	"Golang_RPG/controllers"

	"github.com/astaxie/beego"
)

func init() {
<<<<<<< HEAD
	// TODO: put initializations in a differnt file
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// TODO: add connection string in app.con and unify it on our systems
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("connectionString"))
	orm.RegisterModel(new(models.Users))
	orm.RegisterModel(new(models.Bots))
=======
>>>>>>> 9aaf3168fbb09474dbcf3a15f248267b2e605345
	beego.Router("/api/", &controllers.MainController{})
	beego.Router("/api/register", &controllers.RegisterController{})
	beego.Router("/api/login", &controllers.LoginController{})
	beego.Router("/api/bot", &controllers.BotController{})
	beego.Router("/api/search", &controllers.ShopsSearchController{}) // searches for nearest shops
}
