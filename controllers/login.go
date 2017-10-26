package controllers

import (
	"Golang_RPG/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/astaxie/beego/logs"
)

type LoginController struct {
	beego.Controller
}

type Credentials models.Credentials

func (c *LoginController) Post() {
	x := Credentials{Username: c.GetString("username"), Password: c.GetString("password")}
	o := orm.NewOrm()
	o.Using("default")
	fmt.Println(o.Insert(x))
	fmt.Println("Added new entry to the DB!")
	c.Data["json"] = &x
	// TODO:Hey jude
	// l := logs.GetLogger()
	// l.Println(x)

	c.ServeJSON()
}
