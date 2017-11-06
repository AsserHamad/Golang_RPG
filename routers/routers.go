package routers

import (
	"Golang_RPG/controllers"
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

var checkForAuthorization = func(ctx *context.Context) {

	// ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	// ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	subUrl := strings.Split(ctx.Request.URL.Path, "api/")[1]
	if subUrl == "login" || subUrl == "register" || subUrl == "welcome" {
		return
	}

	fmt.Println("eyo in here")
	_, ok := ctx.Input.Session("id").(int)
	if !ok {
		ctx.ResponseWriter.WriteHeader(401)
		ctx.WriteString("Unauothorized!")
	}
}

func init() {
	beego.InsertFilter("/*", beego.BeforeRouter, checkForAuthorization)
	beego.Router("/api/welcome", &controllers.MainController{})
	beego.Router("/api/register", &controllers.RegisterController{})
	beego.Router("/api/login", &controllers.LoginController{})
	beego.Router("/api/bot", &controllers.BotController{})
	beego.Router("/api/search", &controllers.ShopsSearchController{})                 // searches for nearest shops
	beego.Router("/api/shops/nearestshop", &controllers.NearestShopItemsController{}) // uses result from nearest shop search to display items
	beego.Router("/api/scan", &controllers.ScanController{})
	beego.Router("/api/buyitem", &controllers.BuyItemController{})
}
