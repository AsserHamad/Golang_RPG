package routers

import (
	"Golang_RPG/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/", &controllers.MainController{})
	beego.Router("/api/register", &controllers.RegisterController{})
	beego.Router("/api/login", &controllers.LoginController{})
	beego.Router("/api/bot", &controllers.BotController{})
	beego.Router("/api/search", &controllers.ShopsSearchController{})                 // searches for nearest shops
	beego.Router("/api/shops/nearestshop", &controllers.NearestShopItemsController{}) // uses result from nearest shop search to display items
	beego.Router("/api/scan", &controllers.ScanController{})
	beego.Router("/api/buy/item/:name", &controllers.BuyItemController{})
}
