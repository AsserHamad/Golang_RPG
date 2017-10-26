package main

import (
	"Golang_RPG/models"
	_ "Golang_RPG/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Credentials models.Credentials

func main() {
	orm.RegisterModel(new(Credentials))
	beego.Run()
}
