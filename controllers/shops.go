package controllers

import (
	"Golang_RPG/errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/astaxie/beego"
)

type ShopsController struct {
	beego.Controller
}

func (c *ShopsController) Get(latitude int, longtitude int) {
	resp, err := http.Get("https://maps.googleapis.com/maps/api/place/nearbysearch/json?" + "location=" + strconv.Itoa(latitude) + "," + strconv.Itoa(longtitude) + "&radius=500" + "&key=" + beego.AppConfig.String("gooleKey"))
	if err == nil {
		c.Data["json"] = &errors.InvalidParameters.Message
		c.Ctx.ResponseWriter.WriteHeader(errors.InvalidParameters.HTTPStatus)
	} else {
		fmt.Println(resp)
		c.Data["json"] = &errors.InvalidParameters.Message

	}
}
