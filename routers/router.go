package routers

import (
	"Golang_RPG/controllers"
	"Golang_RPG/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:Pupper1996@/orm_test?charset=utf8")
	orm.RegisterModel(new(models.Credentials))
	beego.Router("/api/", &controllers.MainController{})
	beego.Router("/api/login", &controllers.LoginController{})
}
