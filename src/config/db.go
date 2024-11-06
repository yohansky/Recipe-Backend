package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	// url := os.Getenv(`postgres://PGUSER:PGPASSWORD@PGHOST:PGPORT/PGDATABASE?sslmode=require`)
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		fmt.Println("DATABASE_URL is not set in the environment")
	}
	DB, err = gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		panic("Failed connect to Database")
	} else {
		fmt.Println("Connected to Database")
	}
}
