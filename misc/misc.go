package misc

import (
	"fmt"
	"html"
	"log"
	"main/database"
	"main/services"
)

func Mention(userID int64, fullName string) string {
	return fmt.Sprintf(`<a href="tg://user?id=%d">%s</a>`, userID, html.EscapeString(fullName))
}

func LoadDepartments() {
	log.Println("Loading departments...")
	departments, err := services.GetDepartments()
	if err != nil {
		log.Println("Error fetching departments:", err)
		return
	}

	log.Println("Fetched", len(departments), "departments.")
	for _, dept := range departments {
		log.Printf("Department: %s (Code: %s)", dept.DipDes, dept.Cod)
	}

	if err := database.SaveDepartments(departments); err != nil {
		log.Println("Error saving departments:", err)
		return
	}

	log.Println("Finished loading departments.")
}

func LoadCourses() {
	log.Println("Loading courses...")
	currentPage := 0
	for {
		log.Println("Fetching page:", currentPage)
		courses, err := services.GetCourses(currentPage)
		if err != nil {
			log.Println("Error fetching courses:", err)
			return
		}

		if len(courses.Items) == 0 {
			break
		}

		if err := database.SaveCourses(courses.Items); err != nil {
			log.Println("Error saving courses:", err)
			return
		}

		currentPage++
	}

	log.Println("Finished loading courses.")
}
