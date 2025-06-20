package service

import (
	model "Go_Watchlist/userAuthentication/models"
	repo "Go_Watchlist/userAuthentication/repos"
	// "regexp"
	"errors"
)

// func RegisterUser(req model.AddUserRequest) error {

// 	exists, err := repo.UserExists(req.PAN)
// 	if err != nil {
// 		return err
// 	}
// 	if exists {
// 		return errors.New("User already exists")
// 	}

// 	dbUser := model.AddUserDB{
// 		Name:   req.Name,
// 		PAN:    req.PAN,
// 		Mobile: req.Mobile,
// 		Age:    req.Age,
// 		DOB:    req.DOB,
// 	}
// 	return repo.InsertUser(dbUser)
// }


func RemoveRegisterUser(req model.RemoveUserRequest) error {
	exists, err := repo.UserExistsByPANAndMobile(req.PAN, req.Mobile)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("user with given PAN and mobile does not exist or mismatch")
	}
	return repo.RemoveUserByPAN(req.PAN)
}


