package database

import (
	"main/config"
	"main/models"
)

func SaveDepartments(departments []models.Department) error {
	for _, dept := range departments {
		if err := config.DB.FirstOrCreate(&dept).Error; err != nil {
			return err
		}
	}
	return nil
}

func GetDepartments() ([]models.Department, error) {
	var departments []models.Department
	if err := config.DB.Find(&departments).Error; err != nil {
		return nil, err
	}
	return departments, nil
}
