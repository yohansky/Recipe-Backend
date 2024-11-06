package models

import "gorm.io/gorm"

type Liked_Recipe struct {
	Id       uint   `json:"id"`
	UserId   uint   `json:"user_id"`
	RecipeId uint   `json:"recipe_id"`
	Recipe   Recipe `gorm:"foreignKey:RecipeId"`
}

func (liker *Liked_Recipe) Count(db *gorm.DB) int64 {
	var total int64

	db.Model(&Liked_Recipe{}).Count(&total)

	return total
}

func (liker *Liked_Recipe) Take(db *gorm.DB, limit int, offset int) interface{} {
	var likes []Liked_Recipe

	db.Offset(offset).Limit(limit).Find(&likes)

	return likes
}
