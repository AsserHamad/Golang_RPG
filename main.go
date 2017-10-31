package main

import (
	"Golang_RPG/models"
	_ "Golang_RPG/routers"

	"github.com/astaxie/beego"
)

type Credentials models.Credentials

func main() {
	beego.Run()
}
