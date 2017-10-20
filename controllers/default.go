package controllers

import (
	"fmt"
	"math/rand"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
)

//MainController controls the root API functions
type MainController struct {
	beego.Controller
}

//TODO: Add an extra 'options' component which controls
//			the message options that the user can choose from

//Welcome is a welcoming struct
type Welcome struct {
	ServerStatus   bool     `json:"serverStatus"`
	WelcomeMessage string   `json:"welcomeMessage"`
	Options        []string `json:"options"`
}

//Get gets
func (c *MainController) Get() {
	if c.GetSession("userId") == nil {
		options := []string{"Yes", "No"}
		x := Welcome{true, "Welcome adventurer! Did you happen to visit this realm before?", options}
		// userID := session.Get("UserID")

		c.SetSession("userId", rand.Intn(1000))
		// }

		c.Data["json"] = &x
		// // TODO:Hey jude don't forget to uncomment
		// // l := logs.GetLogger()
		// // l.Println(x.ServerStatus)

	} else {
		options := []string{"Continue"}
		l := logs.GetLogger()
		l.Println(c.GetSession("userId"))
		y := fmt.Sprintf("Welcome, %d", c.GetSession("userId").(int))
		x := Welcome{true, y, options}
		c.Data["json"] = &x
	}

	c.ServeJSON()
}
