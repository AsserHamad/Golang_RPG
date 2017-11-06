package controllers

import (
	"Golang_RPG/models"
)

type win struct {
	Message string
	Enemy   models.Enemies
	Fakka   int
}

func Win(c *ChatController) {
	enemy := c.GetSession("enemy").(models.Enemies)
	c.SetSession("inBattle", false)
	c.Data["json"] = &win{Message: "You won!", Enemy: enemy, Fakka: enemy.Fakka}
	c.ServeJSON()
}
