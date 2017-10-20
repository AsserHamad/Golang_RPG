package controllers

import (
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/logs"
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
	println("I'm a little teapot short and stout")
	options := []string{"Yes", "No"}
	x := Welcome{true, "Welcome adventurer! Did you happen to visit this realm before?", options}
	c.Data["json"] = &x
	// TODO:Hey jude don't forget to uncomment
	// l := logs.GetLogger()
	// l.Println(x.ServerStatus)

	c.ServeJSON()
}
