package services

import (
	"auth/config"
	"auth/models"
	"errors"
)

func CreateUser(user *models.User) error {
	db, err := config.GetDB()

	if err != nil {
		return err
	}

	// Check if username already exists
	var existingUser models.User
	err = db.Where("user_name = ?", user.UserName).First(&existingUser).Error

	if err == nil {
		return errors.New("username alreay exists")
	}

	// Check if email already exists
	err = db.Where("email = ?", user.Email).First(&existingUser).Error

	if err == nil {
		return errors.New("email alreay exists")
	}

	return db.Create(user).Error
}

func FindByUsername(username string) (*models.User, error) {
	db, err := config.GetDB()

	if err != nil {
		return nil, err
	}

	var user models.User
	err = db.Where("user_name = ?", username).First(&user).Error
	return &user, err
}

func FindByEmail(email string) (*models.User, error) {
	db, err := config.GetDB()

	if err != nil {
		return nil, err
	}

	var user models.User
	err = db.Where("email = ?", email).First(&user).Error
	return &user, err
}
