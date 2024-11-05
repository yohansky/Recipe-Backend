package models

type Recipe struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Ingredients string `json:"ingredients"`
	Photo       string `json:"photo"`
	UserId      string `json:"user_id"`
}
