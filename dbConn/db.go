// package databaseConn

// import (
// 	"database/sql"
// 	"fmt"

// 	_ "github.com/lib/pq"
// )

// var DB *sql.DB

//	func Init() {
//		connStr := "host=localhost port=5432 user=postgres password=Pass@123 dbname=postgres sslmode=disable"
//		var err error
//		DB, err = sql.Open("postgres", connStr)
//		if err != nil {
//			panic(err)
//		}
//		err = DB.Ping()
//		if err != nil {
//			panic(err)
//		}
//		fmt.Println("Connected to DB")
//	}
package dbConn

import (
	"Go_Watchlist/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := "host=localhost user=postgres password=Pass@123 dbname=postgres port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	fmt.Println("Database connected successfully")

	err = DB.AutoMigrate(&model.User{}, &model.WatchlistItem{})
	if err != nil {
		log.Fatal("AutoMigrate failed:", err)
	}
}
