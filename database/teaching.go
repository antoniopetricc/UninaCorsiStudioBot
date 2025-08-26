package database

import (
	"main/config"
	"main/models"
)

func SaveTeachings(teachings []models.Teaching) error {
	for _, teaching := range teachings {
		if err := config.DB.Save(&teaching).Error; err != nil {
			return err
		}
	}
	return nil
}

func GetTeachings(codCorso string) ([]models.Teaching, error) {
	var teachings []models.Teaching
	if err := config.DB.Where("codCorso = ?", codCorso).Find(&teachings).Error; err != nil {
		return nil, err
	}
	return teachings, nil
}
