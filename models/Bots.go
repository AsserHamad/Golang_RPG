package models

type Bots struct {
	Id         int    `orm:"auto"`
	Name       string `orm:"size(45)" ,json:"name"`
	Race       string `orm:"size(45)" ,json:"race"`
	Level      int    `orm:"size(45)" ,json:"level"`
	User_id    int    `orm:"size(45)" ,json:"user_id"`
	Experience int    `orm:"size(45)" ,json:"experience"`
	Attack     int    `orm:"size(45)" ,json:"attack"`
	Defense    int    `orm:"size(45)" ,json:"defense"`
	Next_level int    `orm:"size(45)" ,json:"next_level"`
	Fakka      int    `orm:"size(45)" ,json:"fakka"`
}
