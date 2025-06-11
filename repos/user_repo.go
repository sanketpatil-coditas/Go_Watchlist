// package repo

// import (
// 	db "Go_Watchlist/dbConn"
// 	model "Go_Watchlist/models"
// 	"fmt"
// )

//	func InsertUser(u model.User) error {
//		query := `INSERT INTO users (name, pan, mobile, age, dob) VALUES ($1, $2, $3, $4, $5)`
//		_, err := db.DB.Exec(query, u.Name, u.PAN, u.Mobile, u.Age, u.DOB)
//		return err
//	}
//
//	func UserExists(pan string) (bool, error) {
//		query := `SELECT COUNT(*) FROM users WHERE pan = $1`
//		fmt.Println("query", query)
//		var count int
//		err := db.DB.QueryRow(query, pan).Scan(&count)
//		if err != nil {
//			return false, err
//		}
//		return count > 0, nil
//	}
package repo

import (
	"Go_Watchlist/dbConn"
	model "Go_Watchlist/models"
)

func InsertUser(u model.User) error {
	return dbConn.DB.Create(&u).Error
}

func UserExists(pan string) (bool, error) {
	var count int64
	err := dbConn.DB.Model(&model.User{}).Where("pan = ?", pan).Count(&count).Error
	return count > 0, err
}

func RemoveUser(item model.User) error {
	return dbConn.DB.Where("pan = ?", item.PAN).Delete(&model.User{}).Error
	// return dbConn.DB.Delete(&model.User{}).Error
}
