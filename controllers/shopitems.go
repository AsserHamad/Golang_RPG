package controllers

import (
	"Golang_RPG/errors"
	"Golang_RPG/models"
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type NearestShopItemsController struct {
	beego.Controller
}

func (c *NearestShopItemsController) Get() {
	nearestShop := c.GetSession("nearShop")
	if nearestShop == nil {
		c.Data["json"] = &errors.SearchForShop.Message
		c.Ctx.ResponseWriter.WriteHeader(errors.SearchForShop.HTTPStatus)
	} else {
		o := orm.NewOrm()
		var shopItems []*models.Shop_Items
		_, err := o.Raw("SELECT * FROM shop_items WHERE location_id = ?", nearestShop).QueryRows(&shopItems)
		fmt.Println(shopItems[0].Item.Name)
		if err != nil {
			fmt.Println(err.Error())
		}
		var stringItems []string
		stringItems = append(stringItems, "[")
		for _, x := range shopItems {
			stringItems = append(stringItems, fmt.Sprintf(
				"{Name: %s, Required Level: %d, Description: %s, Race: %s, Price: %d},",
				x.Item.Name, x.Item.Required_level, x.Item.Description, x.Item.Race, x.Price))
		}
		if len(stringItems) > 1 {
			stringItems[len(stringItems)-1] = "]"
		} else {
			stringItems = append(stringItems, "]")
		}
		c.Data["json"] = &Response{Message: "Available Items: " + strings.Join(stringItems, "")}
	}
	c.ServeJSON(true)
}
