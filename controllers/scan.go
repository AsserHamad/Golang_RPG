package controllers

import (
	"Golang_RPG/controllers/scan"
	"Golang_RPG/errors"
	"Golang_RPG/models"
	"fmt"
	"math/rand"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/astaxie/beego/logs"
)

type ScanController struct {
	beego.Controller
}

func (c *ScanController) Post() {
	o := orm.NewOrm()
	if c.GetSession("id") != nil && c.GetSession("bot") != nil {
		if c.GetSession("inBattle") == false {
			rand.Seed(time.Now().UTC().UnixNano())
			random := rand.Intn(100)
			fmt.Println("Random number is ", random)
			//50 % chance of finding nothing
			if random < 50 {
				c.Data["json"] = scan.Nothing
			} else {
				//40% for a BATTLE!
				if random < 90 {
					//Battle stuff
					c.SetSession("inBattle", true)
					if c.GetSession("inLocation") == true {
						//Found a BOSS enemy
						var enemies []orm.Params
						_, err := o.Raw("SELECT * FROM enemies WHERE type = ? order by rand() limit 1", "2").Values(&enemies)
						if err != nil {
							c.Data["json"] = &errors.ErrorMessage{Message: err.Error()}
						} else {
							fmt.Println(enemies[0])
							enemy := models.TurnToEnemy(enemies[0])
							c.Data["json"] = scan.EnterBattle("found random BOSS enemy woo", "boss", enemy)
						}
					} else {
						//Found a normal enemy
						var enemies []orm.Params
						_, err := o.Raw("SELECT * FROM enemies WHERE type = ? order by rand() limit 1", "1").Values(&enemies)
						if err != nil {
							c.Data["json"] = &errors.ErrorMessage{Message: err.Error()}
						} else {
							fmt.Println(enemies[0])
							enemy := models.TurnToEnemy(enemies[0])
							c.Data["json"] = scan.EnterBattle("found random enemy woo", "normal", enemy)
						}
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
