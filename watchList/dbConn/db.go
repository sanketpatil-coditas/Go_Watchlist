package db

import (
	model "Go_Watchlist/watchlist/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	DatabaseCreds := "host=localhost user=postgres password=Pass@123 dbname=postgres port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(DatabaseCreds), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	fmt.Println("Database connected successfully")

	err = DB.AutoMigrate(&model.AddWatchListItem{}, &model.CreateWatchlist{})
	if err != nil {
		log.Fatal("AutoMigrate failed:", err)
	}
}
