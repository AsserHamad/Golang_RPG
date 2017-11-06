package controllers

import (
	"Golang_RPG/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type myError struct {
	Message error `json: "message"`
}

type RegisterController struct {
	beego.Controller
}

func ChatRegister(username string, password string, name string, age int, c *ChatController) {

	o := orm.NewOrm()
	x := models.Users{
		Username: username,
		Password: password,
		Name:     name,
		Age:      age,
	}
	_, err := o.Insert(&x)
	if err != nil {
		fmt.Println(err)
		c.Data["json"] = &myError{Message: err}
	} else {
		fmt.Println("Added new entry to the DB!")
		c.SetSession("id", x.Id)
		id := c.GetSession("id")
		fmt.Println("Your new ID iiiiiiiiiis")
		fmt.Println(id)
		//TODO: remove this JSON reply
		c.Data["json"] = &x
	}
	c.ServeJSON()
}
