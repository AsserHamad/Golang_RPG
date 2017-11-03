package battle

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/astaxie/beego/logs"
)

type BattleAttackController struct {
	beego.Controller
}

type we struct {
	Message string `json: message`
}

func (c *BattleAttackController) Get() {
	c.Data["json"] = we{Message: "attacking!"}
	c.ServeJSON()
}
