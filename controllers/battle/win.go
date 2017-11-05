package battle

import "Golang_RPG/models"

type win struct {
	Message string
	Enemy   models.Enemies
	Fakka   int
}

func Win(c *BattleAttackController) {
	enemy := c.GetSession("enemy").(models.Enemies)
	c.Data["json"] = &win{Message: "You won!", Enemy: enemy, Fakka: enemy.Fakka}
	c.ServeJSON()
}
