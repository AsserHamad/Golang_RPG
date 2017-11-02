package models

type Enemies struct {
	Id        int    `orm:"auto"`
	Name      string `orm:"size(100)" ,json:"name"`
	Type      int    `orm:"size(100)" ,json:"type"`
	Location  string `orm:"size(100)" ,json:"location"`
	Attack    int    `orm:"size(100)" ,json:"attack"`
	Defense   int    `orm:"size(100)" ,json:"defense"`
	Pp        int    `orm:"size(100)" ,json:"pp"`
	Agility   int    `orm:"size(100)" ,json:"agility"`
	Maxhp     int    `orm:"size(100)" ,json:"maxhp"`
	Fakka     int    `orm:"size(100)" ,json:"fakka"`
	Drop_item int    `orm:"size(100)" ,json:"drop_item"`
}
