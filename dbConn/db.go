package db

import (
	"Go_Watchlist/config"
	"Go_Watchlist/dbConfig"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	config.LoadConfig("C:/Users/Coditas/Desktop/Projects/Go_Watchlist/config/postgres.yaml") // Load DB config from YAML

	creds := config.AppConfig.Postgres
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		creds.Host, creds.User, creds.Password, creds.DBName, creds.Port, creds.SSLMode,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	fmt.Println("Database connected successfully")

	err = DB.AutoMigrate(&dbConfig.AddWatchListItem{}, &dbConfig.CreateWatchlist{})
	if err != nil {
		log.Fatal("AutoMigrate failed:", err)
	}
}
