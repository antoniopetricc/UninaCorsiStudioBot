package services

import (
	"main/config"
	"main/models"
)

func UpdateStatus(userID int64, status string) error {
	user := models.User{ID: userID, Status: status}
	return config.DB.Save(&user).Error
}

func GetUserStatus(userID int64) (string, error) {
	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		return "", err
	}
	return user.Status, nil
}
