package models

type Comment struct {
	Id       uint   `json:"id"`
	Text     string `json:"text"`
	UserId   uint   `json:"user_id"`
	User     User   `gorm:"foreignKey:UserId" json:"user"`
	RecipeId uint   `json:"recipe_id"`
}
