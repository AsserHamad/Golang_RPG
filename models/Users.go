package models

type Users struct {
	Id       int    `orm:"auto"`
	Username string `orm:"size(100)" ,json:"username"`
	Password string `orm:"size(100)" ,json:"password"`
	Name     string `orm:"size(100)" ,json:"name"`
	Age      string `orm:"size(100)" ,json:"age"`
}
