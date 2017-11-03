package battle

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/astaxie/beego/logs"
)

type BattleDefendController struct {
	beego.Controller
}

func (c *BattleDefendController) Get() {
	c.Data["json"] = we{Message: "defending!"}
	c.ServeJSON()
}
