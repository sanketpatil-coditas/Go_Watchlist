package repo

import (
	"Go_Watchlist/userAuthentication/db"
	"Go_Watchlist/userAuthentication/models"
)

func InsertUser(u models.AddUserDB) error {
	return db.DB.Create(&u).Error
}

func UserExists(pan string) (bool, error) {
	var count int64
	err := db.DB.Model(&models.AddUserDB{}).Where("pan = ?", pan).Count(&count).Error
	return count > 0, err
}

// func RemoveUserByPAN(pan string) error {
// 	return db.DB.Where("pan = ?", pan).Delete(&model.AddUserDB{}).Error
// }
