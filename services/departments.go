package services

import (
	"encoding/json"
	"fmt"
	"main/models"
	"net/http"
)

func GetDepartments() ([]models.Department, error) {
	url := "https://www.corsi.unina.it/corsidistudio-be/v1/departments"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get departments: %s", resp.Status)
	}

	var result []models.Department
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}
