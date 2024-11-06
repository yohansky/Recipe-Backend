package helper

import (
	"backend-recipe/src/config"
	"backend-recipe/src/models"
)

func Migrate() {
	config.DB.AutoMigrate(
		&models.User{},
		&models.Comment{},
		&models.Liked_Recipe{},
		&models.Recipe{},
		&models.Saved_Recipe{},
	)
}
