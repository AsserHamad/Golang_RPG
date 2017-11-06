package controllers

import (
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

type ChatController struct {
	beego.Controller
}

func (c *ChatController) Post() {
	message := strings.Split(c.GetString("message"), " ")
	switch message[0] {
	case "login":
		ChatLogin(message[1], message[2], c)
	case "register":
		age, _ := strconv.Atoi(strings.Split(message[4], "\n")[0])
		ChatRegister(message[1], message[2], message[3], age, c)
	case "scan":
		ChatScan(c)
	case "attack":
		ChatAttack(c)
	case "defend":
		ChatDefend(c)
	case "location":
	}
}
