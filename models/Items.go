package models

type Items struct {
	Id            int `orm:"auto"`
	RequiredLevel int `orm:"size(45)" ,json:"required_level"`
	Name          string
	Description   string
	Race          string
	Price         int `orm:"size(45)" ,json:"price"`
}
