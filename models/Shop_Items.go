package models

type Shop_Items struct {
	Id          int `orm:"auto"`
	Location_Id int
	Item        *Items `orm:"rel(one)"`
	Price       int
}
