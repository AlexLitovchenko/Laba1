package db

import (
	"lab1/go-rest-api/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func InitDB() error {
	db, err := gorm.Open(postgres.Open(config.GetConfig().Db), &gorm.Config{})
	database = db
	return err
}

func GetDb() *gorm.DB {
	return database
}
