package repo

import (
	"Go_Watchlist/dbConn"
	// "Go_Watchlist/userAuthentication/models"
	"Go_Watchlist/dbConfig"
)

// func InsertUser(u model.AddUserDB) error {
// 	return db.DB.Create(&u).Error
// }

// func UserExists(pan string) (bool, error) {
// 	var count int64
// 	err := db.DB.Model(&model.AddUserDB{}).Where("pan = ?", pan).Count(&count).Error
// 	return count > 0, err
// }
func UserExistsByPANAndMobile(pan string, mobile string) (bool, error) {
	var count int64
	err := db.DB.Model(&dbConfig.AddUserDB{}).
		Where("pan = ? AND mobile = ?", pan, mobile).
		Count(&count).Error
	return count > 0, err
}


func RemoveUserByPAN(pan string) error {
	return db.DB.Where("pan = ?", pan).Delete(&dbConfig.AddUserDB{}).Error
}

