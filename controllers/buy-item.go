package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type BuyItemController struct {
	beego.Controller
}

func (c *BuyItemController) Get() {
	var itemName string = c.GetString("name")
	o := orm.NewOrm()

}
