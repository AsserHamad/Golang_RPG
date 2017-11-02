package models

type Locations struct {
	Id        int     `orm:"auto"`
	Latitude  float64 `orm:"size(45)" ,json:"latitude"`
	Longitude float64 `orm:"size(45)" ,json:"longitude"`
	Name      string  `orm:"size(45)" ,json:"name"`
	Type      string  `orm:"size(45)" ,json:"type"`
}
