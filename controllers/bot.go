package controllers

import (
	"Golang_RPG/errors"
	"Golang_RPG/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type BotController struct {
	beego.Controller
}
type Ayhabal struct {
	Id int `json:"id"`
}

func (c *BotController) Get() {
	o := orm.NewOrm()
	_id := c.GetSession("id")
	if _id != nil {
		id := _id.(int)
		bot := models.Bots{User_id: id}
		err := o.Read(&bot)
		if err != nil {
			c.Data["json"] = &errors.ErrorMessage{Message: err.Error()}
			c.Ctx.ResponseWriter.WriteHeader(401)
		} else {
			c.Data["json"] = &bot
		}
	} else {
		c.Data["json"] = &errors.NotLoggedIn.Message
		c.Ctx.ResponseWriter.WriteHeader(401)
	}
	c.ServeJSON()
}

func (c *BotController) Post() {

	o := orm.NewOrm()
	_id := c.GetSession("id")
	if _id != nil {
		id := _id.(int)
		_bot := models.Bots{User_id: id}
		err := o.Read(&_bot)
		if err != nil {
			bot := models.Bots{
				Name:    c.GetString("name"),
				Race:    c.GetString("race"),
				Level:   1,
				User_id: id,
				Attack:  10,
				Defense: 10,
				Fakka:   10,
			}
			_, err := o.Insert(&bot)
			if err != nil {
				c.Data["json"] = &errors.Err{Message: err}
				c.Ctx.ResponseWriter.WriteHeader(401)
			} else {
				c.Data["json"] = &bot
			}
		} else {
			c.Data["json"] = &errors.HaveBot.Message
			c.Ctx.ResponseWriter.WriteHeader(401)
		}
	} else {
		c.Data["json"] = &errors.NotLoggedIn.Message
		c.Ctx.ResponseWriter.WriteHeader(401)
	}
	c.ServeJSON()
}
