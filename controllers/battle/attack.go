package battle

import (
	"Golang_RPG/models"
	"fmt"
	"math/rand"
	"time"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/astaxie/beego/logs"
)

type BattleAttackController struct {
	beego.Controller
}

type we struct {
	Message string         `json: message`
	Enemy   models.Enemies `json:"enemies"`
}

func attack(c *BattleAttackController) {
	player := c.GetSession("bot").(models.Bots)
	enemy := c.GetSession("enemy").(models.Enemies)

	enemyCurrentHealth, _ := c.GetSession("enemyCurrentHealth").(int)

	rand.Seed(time.Now().UTC().UnixNano())
	// random := rand.Intn(500)
	// if random <= enemy.Agility {
	// 	random = 0
	// } else {
	// 	random = 1
	// }
	random := rand.Intn(enemy.Agility)

	enemyCurrentHealth =
		enemyCurrentHealth - (player.Attack-enemy.Defense/100)*random
	fmt.Println("Enemy health ", enemyCurrentHealth, "(calcuation)", (player.Attack-enemy.Defense/100)*random)

	if enemyCurrentHealth <= 0 {
		Win(c)
	} else {
		c.SetSession("enemyCurrentHealth", enemyCurrentHealth)
		EnemyTurn(c, enemy, player)
	}
}

func (c *BattleAttackController) Post() {
	attack(c)
}
