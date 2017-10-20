package controllers

import (
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/logs"
)

type LoginController struct {
	beego.Controller
}

type Credentials struct {
	username string
	password string
}

func (c *LoginController) Post() {
	username := c.Ctx.Input.Request.Form.Get("username")
	password := c.Ctx.Input.Request.Form.Get("password")

	c.Data["json"] = &x
	// TODO:Hey jude
	// l := logs.GetLogger()
	// l.Println(x.ServerStatus)

	c.ServeJSON()
}
