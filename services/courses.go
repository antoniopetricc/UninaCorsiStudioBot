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
