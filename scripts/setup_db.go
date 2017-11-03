package script

import (
	"Golang_RPG/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("connectionString"))
	orm.RegisterModel(new(models.Users))
	orm.RegisterModel(new(models.Bots))
	orm.RegisterModel(new(models.Enemies))
	orm.RegisterModel(new(models.Items))
	orm.RegisterModel(new(models.Locations))
}
