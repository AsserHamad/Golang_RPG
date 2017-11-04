package models

import "github.com/astaxie/beego/orm"

type Items struct {
	Id             int    `orm:"auto"`
	Required_level int    `orm:"size(45)" ,json:"required_level"`
	Name           string `orm:"size(45)" ,json:"name"`
	Description    string `orm:"size(200)" ,json:"description"`
	Race           string `orm:"size(45)" ,json:"race"`
	Price          int    `orm:"size(45)" ,json:"price"`
}

func TurnToItem(enemy orm.Params) Items {
	ID, _ := StrToInt(enemy["id"].(string))
	Required_level, _ := StrToInt(enemy["required_level"].(string))
	Name, _ := enemy["name"].(string)
	Description, _ := enemy["description"].(string)
	Race, _ := enemy["race"].(string)
	Price, _ := StrToInt(enemy["price"].(string))

	return Items{
		Id:             ID,
		Required_level: Required_level,
		Name:           Name,
		Description:    Description,
		Race:           Race,
		Price:          Price,
	}
}
