package battle

type lose struct {
	Message string
}

func Lose(c *BattleAttackController) {
	c.Data["json"] = &lose{Message: "You lose..."}
	c.ServeJSON()

}
