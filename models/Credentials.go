package models

type Credentials struct {
	Id       int
	Username string `json:"username"`
	Password string `json:"password"`
}
