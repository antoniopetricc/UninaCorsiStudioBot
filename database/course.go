package database

import (
	"main/config"
	"main/models"
)

func GetCourse(cod string) (models.Course, error) {
	var course models.Course
	if err := config.DB.Where("cod = ?", cod).First(&course).Error; err != nil {
		return models.Course{}, err
	}
	return course, nil
}

func SaveCourses(courses []models.Course) error {
	for _, course := range courses {
		if err := config.DB.Save(&course).Error; err != nil {
			return err
		}
	}
	return nil
}

func GetCourses(page int) ([]models.Course, int64, error) {
	var corsi []models.Course
	var total int64
	limit := 10
	offset := page * limit
	if err := config.DB.Limit(limit).Offset(offset).Find(&corsi).Error; err != nil {
		return nil, 0, err
	}
	if err := config.DB.Model(&models.Course{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return corsi, total, nil
}

func SearchCourses(query string, page int) ([]models.Course, int64, error) {
	var corsi []models.Course
	var total int64
	limit := 10
	offset := page * limit

	if err := config.DB.Where("nome LIKE ?", "%"+query+"%").Limit(limit).Offset(offset).Find(&corsi).Error; err != nil {
		return nil, 0, err
	}
	if err := config.DB.Model(&models.Course{}).Where("nome LIKE ?", "%"+query+"%").Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return corsi, total, nil
}

func UpdateCourseDescription(cod string, description string) error {
	return config.DB.Model(&models.Course{}).Where("cod = ?", cod).Update("in_breve", description).Error
}
