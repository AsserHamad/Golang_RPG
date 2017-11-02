package models

type Items struct {
	Id             int    `orm:"auto"`
	Required_level int    `orm:"size(45)" ,json:"required_level"`
	Name           string `orm:"size(45)" ,json:"name"`
	Description    string `orm:"size(200)" ,json:"description"`
	Race           string `orm:"size(45)" ,json:"race"`
	Price          int    `orm:"size(45)" ,json:"price"`
}
