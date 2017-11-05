package controllers

import (
	"strings"

	"github.com/astaxie/beego"
)

type ChatController struct {
	beego.Controller
}

func (c *ChatController) Post() {
	message := strings.Split(c.GetString("message"), " ")
	switch message[0] {
	case "bot":
		c.BotPost(message[1], message[2])
	case "register":
		age, _ := StrToInt(message[4])
		c.RegisterPost(message[1], message[2], message[3], age)
	case "login":
		c.LoginPost(message[1], message[2])
	case "scan":
	}
}

func (c *ChatController) Get() {
	message := strings.Split(c.GetString("message"), " ")
	switch message[0] {
	case "attack":
		c.AttackGet()
	case "defend":
	case "item":
	case "default":
	case "shops-search":
		c.ShopsSearchGet()
	case "":
	}
}
