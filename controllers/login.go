package controllers

import (
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/logs"
)

type LoginController struct {
	beego.Controller
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *LoginController) Post() {
	x := Credentials{c.GetString("username"), c.GetString("password")}

	c.Data["json"] = &x
	// TODO:Hey jude
	// l := logs.GetLogger()
	// l.Println(x)

	c.ServeJSON()
}
