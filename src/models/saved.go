package models

import "gorm.io/gorm"

type Saved struct {
	Id       uint   `json:"id"`
	UserId   uint   `json:"user_id"`
	RecipeId uint   `json:"recipe_id"`
	Recipe   Recipe `gorm:"foreignKey:RecipeId"`
}

func (saved *Saved) Count(db *gorm.DB) int64 {
	var total int64

	db.Model(&Saved{}).Count(&total)

	return total
}

func (saved *Saved) Take(db *gorm.DB, limit int, offset int) interface{} {
	var saves []Saved

	db.Offset(offset).Limit(limit).Find(&saves)

	return saves
}
