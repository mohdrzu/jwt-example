package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func GetConnection() (*gorm.DB, error) {
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&User{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
