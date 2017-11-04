package main

import (
	_ "Golang_RPG/routers"
	_ "Golang_RPG/scripts"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func main() {
	orm.Debug = true
	beego.Run()
}
