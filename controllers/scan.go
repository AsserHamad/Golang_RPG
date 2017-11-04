package controllers

import (
	"Golang_RPG/controllers/scan"
	"Golang_RPG/errors"
	"Golang_RPG/models"
	"math/rand"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/astaxie/beego/logs"
)

type ScanController struct {
	beego.Controller
}

func (c *ScanController) Post() {
	if c.GetSession("id") != nil && c.GetSession("bot") != nil {
		if c.GetSession("inBattle") == false {
			random := rand.Intn(100)
			//50 % chance of finding nothing
			if random < 50 {
				c.Data["json"] = scan.Nothing
			} else {
				//40% for a BATTLE!
				if random < 90 {
					//Battle stuff
					if c.GetSession("inLocation") == true {
						c.SetSession("inBattle", true)
						//TODO: Search for a real boss
						c.Data["json"] = scan.EnterBattle("Jojo's Self Confidence", "boss", models.Enemies{})
					} else {
						c.SetSession("inBattle", true)
						c.Data["json"] = scan.EnterBattle("Omar's moustache", "normal", models.Enemies{})
					}
				} else {
					//10% chance of finding an item
					item := scan.FoundItem("You found an item!", models.Items{})
					c.Data["json"] = item
				}
			}

		} else {
			c.Data["json"] = &errors.ErrorMessage{Message: "You're already in a battle!"}
		}
	} else {
		c.Data["json"] = &errors.ErrorMessage{Message: "Please login and/or create a bot"}
	}
	c.ServeJSON()
}
