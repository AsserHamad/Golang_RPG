package controllers

import (
	"Golang_RPG/models"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/astaxie/beego/logs"
)

func StrToInt(str string) (int, error) {
	nonFractionalPart := strings.Split(str, ".")
	return strconv.Atoi(nonFractionalPart[0])
}

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Post() {
	id, _ := StrToInt(c.GetString("id"))
	x := models.Credentials{Id: id, Username: c.GetString("username"), Password: c.GetString("password")}
	o := orm.NewOrm()
	fmt.Println(o.Insert(&x))

	fmt.Println("Added new entry to the DB!")
	c.Data["json"] = &x
	// TODO:Hey jude
	// l := logs.GetLogger()
	// l.Println(x)

	c.ServeJSON()
}
