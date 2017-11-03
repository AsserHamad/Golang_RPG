package routers

import (
	"Golang_RPG/controllers/battle"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/battle/attack", &battle.BattleAttackController{})
	beego.Router("/api/battle/defend", &battle.BattleDefendController{})
}
