package models

type Bots struct {
	Id         int    `orm:"auto"`
	Name       string `orm:"size(45)" ,json:"name"`
	Race       string `orm:"size(45)" ,json:"race"`
	Level      int    `json:"level"`
	User_id    int    `json:"user_id"`
	Experience int    `json:"experience"`
	Attack     int    `json:"attack"`
	Defense    int    `json:"defense"`
	Next_level int    `json:"next_level"`
	Fakka      int    `json:"fakka"`
}
