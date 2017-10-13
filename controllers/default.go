package controllers

import (
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/logs"
)

type Welcome struct {
  ServerStatus bool `json:"serverStatus"`
}

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	x := Welcome{true}
	c.Data["json"] = &x
	// TODO:Hey jude
  // l := logs.GetLogger()
	// l.Println(x.ServerStatus)

	c.ServeJSON();
}
