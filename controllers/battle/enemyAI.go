package battle

import (
	"Golang_RPG/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	// "github.com/astaxie/beego/logs"
)

type Briefing struct {
	MaxHealth      int
	CurrentHealth  int
	EnemyMaxHealth int
	EnemyHealth    int
}

func EnemyTurn(c *BattleAttackController, enemy models.Enemies, player models.Bots) {
	//TODO: Add boss/skill logic
	playerCurrentHealth, _ := c.GetSession("playerCurrentHealth").(int)
	enemyCurrentHealth, _ := c.GetSession("enemyCurrentHealth").(int)

	formula := enemy.Power
	if c.GetSession("playerDefending").(bool) == true {
		formula /= 2
	}

	playerCurrentHealth =
		// playerCurrentHealth - (enemy.Power * (100 - (player.Defense / 500)))
		playerCurrentHealth - formula
	fmt.Println("Your health ", playerCurrentHealth)
	if playerCurrentHealth <= 0 {
		Lose(c)
	} else {
		c.SetSession("playerCurrentHealth", playerCurrentHealth)
		c.SetSession("enemyCurrentHealth", enemyCurrentHealth)
		c.Data["json"] = &Briefing{
			MaxHealth:      player.Maxhp,
			CurrentHealth:  playerCurrentHealth,
			EnemyMaxHealth: enemy.Maxhp,
			EnemyHealth:    enemyCurrentHealth,
		}
		c.ServeJSON()
	}
}
