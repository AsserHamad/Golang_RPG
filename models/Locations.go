package models

type Locations struct {
	Id        int     `orm:"auto"`
	Latitude  float64 `orm:"size(100)" ,json:"latitude"`
	Longitude float64 `orm:"size(100)" ,json:"longitude"`
	Name      string  `orm:"size(100)" ,json:"name"`
	Type      string  `orm:"size(100)" ,json:"type"`
}
