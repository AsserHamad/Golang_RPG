package controllers

import (
	"Golang_RPG/models"
	"strconv"
	"strings"

	"Golang_RPG/errors"

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

type Success struct {
	Message string `json: message`
}

func (c *LoginController) Post() {

	o := orm.NewOrm()
	user := models.Users{Username: c.GetString("username"), Password: c.GetString("password")}
	err := o.Read(&user, "Username", "Password")
	user.Password = ""

	if err == orm.ErrNoRows {
		c.Data["json"] = &errors.WrongCredentials.Message
	} else {
		c.SetSession("userId", user.Id)
		c.Data["json"] = &Success{Message: "Welcome!"}
	}

	c.ServeJSON()
}
