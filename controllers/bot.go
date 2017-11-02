package controllers

import (
	"github.com/astaxie/beego"
)

type BotController struct {
	beego.Controller
}
type Ayhabal struct {
	Id int `json:"id"`
}

func (c *BotController) Get() {
	z := c.GetSession("id")
	if z != nil {
		x := Ayhabal{Id: z.(int)}
		c.Data["json"] = x
	} else {
		x := Ayhabal{Id: -1}
		c.Data["json"] = x
	}
	c.ServeJSON()
}
