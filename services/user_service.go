package service

import (
	model "Go_Watchlist/models"
	repo "Go_Watchlist/repos"
	"errors"
	"regexp"
)

func isValidPAN(pan string) bool {
	return len(pan) == 10
}

func isValidMobile(mobile string) bool {
	matched, _ := regexp.MatchString(`^\d{10}$`, mobile)
	return matched
}
func isFilledRequiredDetails(name string, age int, dob string) bool {
	if name == "" || age == 0 || dob == "" {
		return false
	}
	return true
}

func RegisterUser(u model.User) error {
	if !isValidPAN(u.PAN) {
		return errors.New("PAN must be exactly 10 characters")
	}
	if !isValidMobile(u.Mobile) {
		return errors.New("Mobile number must be exactly 10 digits")
	}
	if !isFilledRequiredDetails(u.Name, u.Age, u.DOB) {
		return errors.New("Please fill required details")
	}
	exists, err := repo.UserExists(u.PAN)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("User already exists")
	}
	return repo.InsertUser(u)
}
func RemoveRegisterUser(item model.User) error {
	return repo.RemoveUser(item)
}
