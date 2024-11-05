package models

type Saved struct {
	Id       uint   `json:"id"`
	UserId   uint   `json:"user_id"`
	RecipeId uint   `json:"recipe_id"`
	Recipe   Recipe `gorm:"foreignKey:RecipeId"`
}
