package controllers

import (
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/logs"
)

type MainController struct {
	beego.Controller
}

type Welcome struct {
	ServerStatus bool `json:"serverStatus"`
}

func (c *MainController) Get() {
	x := Welcome{true}
	c.Data["json"] = &x
	// TODO:Hey jude
	// l := logs.GetLogger()
	// l.Println(x.ServerStatus)

	c.ServeJSON()
}
