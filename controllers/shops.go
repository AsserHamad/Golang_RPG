package controllers

import (
	"Golang_RPG/errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/astaxie/beego"
)

type ShopsController struct {
	beego.Controller
}

func (c *ShopsController) Get() {
	latitude, _ := c.GetInt("latitude")
	longtitude, _ := c.GetInt("longtitude")
	resp, err := http.Get("https://maps.googleapis.com/maps/api/place/nearbysearch/json?" + "location=" + strconv.Itoa(latitude) + "," + strconv.Itoa(longtitude) + "&radius=500" + "&key=" + beego.AppConfig.String("googleKey"))
	if err != nil {
		c.Data["json"] = &errors.InvalidParameters.Message
		c.Ctx.ResponseWriter.WriteHeader(errors.InvalidParameters.HTTPStatus)
	} else {
		bytes, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(bytes))
		c.Data["json"] = &errors.InvalidParameters.Message

	}
	c.ServeJSON(true)
}
