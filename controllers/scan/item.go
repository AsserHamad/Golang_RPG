package scan

import "Golang_RPG/models"

type Item struct {
	Message string       `json:"message"`
	Item    models.Items `json:"item"`
}

func FoundItem(message string, item models.Items) *Item {
	return &Item{Message: message, Item: item}
}
