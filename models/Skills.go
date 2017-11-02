package models

type Skills struct {
	Id             int    `orm:"auto"`
	Name           string `orm:"size(100)" ,json:"name"`
	Race           string `orm:"size(100)" ,json:"race"`
	Required_level int    `orm:"size(100)" ,json:"required_level"`
}
