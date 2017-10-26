package routers

import (
	"Golang_RPG/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8")
	beego.Router("/api/", &controllers.MainController{})
	beego.Router("/api/login", &controllers.LoginController{})
}
