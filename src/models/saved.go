package models

import "gorm.io/gorm"

type Saved_Recipe struct {
	Id       uint   `json:"id"`
	UserId   uint   `json:"user_id"`
	RecipeId uint   `json:"recipe_id"`
	Recipe   Recipe `gorm:"foreignKey:RecipeId"`
}

func (saved *Saved_Recipe) Count(db *gorm.DB) int64 {
	var total int64

	db.Model(&Saved_Recipe{}).Count(&total)

	return total
}

func (saved *Saved_Recipe) Take(db *gorm.DB, limit int, offset int) interface{} {
	var saves []Saved_Recipe

	db.Offset(offset).Limit(limit).Find(&saves)

	return saves
}
