package controllers

import (
	"Golang_RPG/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Post() {

	o := orm.NewOrm()
	z, _ := StrToInt(c.GetString("age"))
	x := models.RegisterCredentials{
		Username: c.GetString("username"),
		Password: c.GetString("password"),
		Name:     c.GetString("name"),
		Age:      z,
	}
	status, err := o.Insert(&x)
	if err != nil {

	}
	fmt.Println()

	fmt.Println("Added new entry to the DB!")
	c.Data["json"] = &x
	// TODO:Hey jude
	// l := logs.GetLogger()
	//	 l.Println(x)

	c.ServeJSON()
}
