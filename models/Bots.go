package models

type Bots struct {
	Id         int    `orm:"auto"`
	Name       string `orm:"size(100)" ,json:"name"`
	Race       string `orm:"size(100)" ,json:"race"`
	Level      int    `orm:"size(100)" ,json:"level"`
	User_id    int    `orm:"size(100)" ,json:"user_id"`
	Experience int    `orm:"size(100)" ,json:"experience"`
	Attack     int    `orm:"size(100)" ,json:"attack"`
	Defense    int    `orm:"size(100)" ,json:"defense"`
	Next_level int    `orm:"size(100)" ,json:"next_level"`
	Fakka      int    `orm:"size(100)" ,json:"fakka"`
}
