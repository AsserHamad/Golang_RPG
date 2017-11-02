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

func (c *RegisterController) Post() {

	o := orm.NewOrm()
	z, _ := StrToInt(c.GetString("age"))
	x := models.Users{
		Username: c.GetString("username"),
		Password: c.GetString("password"),
		Name:     c.GetString("name"),
		Age:      z,
	}
	status, err := o.Insert(&x)
	if err != nil {
		fmt.Println(err)
		ayesmvariable := myError{Message: err}
		c.Data["json"] = &ayesmvariable
	} else {
		fmt.Println("Added new entry to the DB!")
		fmt.Println(status)
		c.Data["json"] = &x
	}

	c.ServeJSON()
}
