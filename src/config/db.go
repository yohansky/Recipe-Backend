package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	url := os.Getenv(`postgres://PGUSER:PGPASSWORD@PGHOST:PGPORT/PGDATABASE?sslmode=require`)
	var err error
	DB, err = gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		panic("Failed connect to Database")
	} else {
		fmt.Println("Connected to Database")
	}
}
