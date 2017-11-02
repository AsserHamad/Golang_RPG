package controllers

import (
	"Golang_RPG/errors"
	"Golang_RPG/models"
	"context"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"googlemaps.github.io/maps"
)

type ShopsController struct {
	beego.Controller
}

type Response struct {
	Message string
}

func (c *ShopsController) Get() {
	latitude, _ := c.GetFloat("latitude")
	longtitude, _ := c.GetFloat("longtitude")
	c2, err := maps.NewClient(maps.WithAPIKey(beego.AppConfig.String("googleKey")))
	r := &maps.NearbySearchRequest{Location: &maps.LatLng{Lat: latitude, Lng: longtitude}, RankBy: "distance", Type: "stadium"}
	resp, err := c2.NearbySearch(context.Background(), r)
	if err != nil {
		c.Data["json"] = &errors.InvalidParameters.Message
		fmt.Println(err.Error())
		c.Ctx.ResponseWriter.WriteHeader(errors.InvalidParameters.HTTPStatus)
	} else {
		nearestLocation := resp.Results[0]
		o := orm.NewOrm()
		user := models.Locations{Name: nearestLocation.Name}
		err := o.Read(&user, "Name")
		if err == orm.ErrNoRows {
			c.Data["json"] = &Response{Message: "No nearby shops"}
		} else {
			var (
				latitudeOfNearest   string = strconv.FormatFloat(nearestLocation.Geometry.Location.Lat, 'f', -1, 64)
				longtitudeOfNearest string = strconv.FormatFloat(nearestLocation.Geometry.Location.Lng, 'f', -1, 64)
				destination         string = latitudeOfNearest + "," + longtitudeOfNearest
				origin              string = c.GetString("latitude") + "," + c.GetString("longtitude")
			)

			r := &maps.DistanceMatrixRequest{Origins: []string{origin}, Destinations: []string{destination}}
			resp, err := c2.DistanceMatrix(context.Background(), r)
			if err != nil {
				c.Data["json"] = &errors.InvalidParameters.Message
				fmt.Println(err.Error())
				c.Ctx.ResponseWriter.WriteHeader(errors.InvalidParameters.HTTPStatus)
			} else {
				fmt.Println(resp.Rows[0].Elements[0].Distance)
				c.Data["json"] = &Response{Message: "hi"}
			}
		}
		c.ServeJSON(true)
	}
}
