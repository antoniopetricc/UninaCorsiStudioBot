package services

import (
	"encoding/json"
	"fmt"
	"main/models"

	"net/http"
)

func GetCourses(page int) (models.CoursesResponse, error) {
	url := fmt.Sprintf("https://www.corsi.unina.it/corsidistudio-be/v1/courses?page=%d&size=10", page)

	resp, err := http.Get(url)
	if err != nil {
		return models.CoursesResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.CoursesResponse{}, fmt.Errorf("failed to get courses: %s", resp.Status)
	}

	var result models.CoursesResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return models.CoursesResponse{}, err
	}

	return result, nil
}

func GetCourseDescription(cod string) (models.CourseDescriptionResponse, error) {
	url := fmt.Sprintf("https://www.corsi.unina.it/corsidistudio-be/v1/courses/details/%s/sua/latest/description", cod)

	resp, err := http.Get(url)
	if err != nil {
		return models.CourseDescriptionResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.CourseDescriptionResponse{}, fmt.Errorf("failed to get course description: %s", resp.Status)
	}

	var result models.CourseDescriptionResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return models.CourseDescriptionResponse{}, err
	}

	return result, nil
}

func GetCourseTeaching(cod string, page int) (models.TeachingsResponse, error) {
	url := fmt.Sprintf("https://www.corsi.unina.it/corsidistudio-be/v1/courses/details/%s/teachings/2025?page=%d&size=10", cod, page)

	resp, err := http.Get(url)
	if err != nil {
		return models.TeachingsResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.TeachingsResponse{}, fmt.Errorf("failed to get course teaching: %s", resp.Status)
	}

	var result models.TeachingsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return models.TeachingsResponse{}, err
	}

	return result, nil
}
