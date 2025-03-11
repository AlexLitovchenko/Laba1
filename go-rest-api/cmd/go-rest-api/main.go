package main

import (
	"lab1/go-rest-api/internal/config"
	"lab1/go-rest-api/internal/db"
	"lab1/go-rest-api/internal/handler"
	"lab1/go-rest-api/internal/models"
	"log"
)

func main() {
	config.InitConfig()
	if err := db.InitDB(); err != nil {
		panic(err)
	}

	db.GetDb().AutoMigrate(
		&models.AccessToken{}, &models.Customer{},
		&models.OrderCart{}, &models.OrderCartItem{},
		&models.Product{},
	)

	router := handler.SetupRoutes()

	log.Println("Server is running on port 8080")
	router.Run(":8080")
}
