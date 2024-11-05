package helper

import (
	"backend-recipe/src/config"
	"backend-recipe/src/models"
)

func Migrate() {
	config.DB.AutoMigrate(
		&models.User{},
		&models.Comment{},
		&models.Liker{},
		&models.Recipe{},
		&models.Saved{},
	)
}
