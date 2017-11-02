package models

type Enemies struct {
	Id        int    `orm:"auto"`
	Name      string `orm:"size(45)" ,json:"name"`
	Type      int    `orm:"size(45)" ,json:"type"`
	Location  string `orm:"size(45)" ,json:"location"`
	Attack    int    `orm:"size(45)" ,json:"attack"`
	Defense   int    `orm:"size(45)" ,json:"defense"`
	Pp        int    `orm:"size(45)" ,json:"pp"`
	Agility   int    `orm:"size(45)" ,json:"agility"`
	Maxhp     int    `orm:"size(45)" ,json:"maxhp"`
	Fakka     int    `orm:"size(45)" ,json:"fakka"`
	Drop_item int    `orm:"size(45)" ,json:"drop_item"`
}
