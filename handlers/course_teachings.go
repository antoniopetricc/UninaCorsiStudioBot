package handlers

import (
	"fmt"
	"log"
	"main/database"
	"main/models"
	"main/services"
	"strconv"
	"strings"

	"github.com/GoBotApiOfficial/gobotapi"
	"github.com/GoBotApiOfficial/gobotapi/methods"
	"github.com/GoBotApiOfficial/gobotapi/types"
)

func buildCourseTeachingsReplyMarkup(course models.Course, currentPage int) *types.InlineKeyboardMarkup {
	return &types.InlineKeyboardMarkup{
		InlineKeyboard: [][]types.InlineKeyboardButton{
			{
				{
					Text:         "🔙 Torna indietro",
					CallbackData: fmt.Sprintf("course_info_%s_%d", course.Cod, currentPage),
				},
			},
		},
	}
}

func CourseTeachings(client *gobotapi.Client, update types.CallbackQuery) {

	if strings.HasPrefix(update.Data, "course_teachings_") {

		parts := strings.SplitN(update.Data, "_", 4)
		courseCod := parts[2]
		currentPage, err := strconv.Atoi(parts[3])
		if err != nil {
			currentPage = 0
		}

		course, err := database.GetCourse(courseCod)
		if err != nil {
			log.Printf("Error fetching course: %v", err)
			return
		}

		teachings, err := database.GetTeachings(courseCod)
		if err != nil {
			log.Printf("Error fetching teachings from DB: %v", err)
		}

		if len(teachings) == 0 {
			teachingsPage := 0
			for {
				teachingsResponse, err := services.GetCourseTeaching(courseCod, teachingsPage)
				if err != nil {
					log.Printf("Error fetching course teachings from API: %v", err)
					break
				}

				if len(teachingsResponse.Items) == 0 {
					break
				}

				teachings = append(teachings, teachingsResponse.Items...)

				if err := database.SaveTeachings(teachingsResponse.Items); err != nil {
					log.Printf("Error saving teachings: %v", err)
				}

				if teachingsResponse.CurrentPage >= teachingsResponse.TotalPages-1 {
					break
				}
				teachingsPage++
			}
		}

		messageText := fmt.Sprintf("🎒 Insegnamenti (%d)", len(teachings))
		if len(teachings) == 0 {
			messageText += "\n\nNessun insegnamento trovato per questo corso."
		} else {
			messageText += "\n\n"
			for _, teaching := range teachings {
				messageText += fmt.Sprintf("• <b>%s</b>: %s\n", teaching.CodInsegnamento, teaching.DescInsegnamento)
			}
		}

		_, err = client.Invoke(&methods.EditMessageText{
			ChatID:      update.Message.Chat.ID,
			MessageID:   update.Message.MessageID,
			Text:        messageText,
			ParseMode:   "HTML",
			ReplyMarkup: buildCourseTeachingsReplyMarkup(course, currentPage),
		})

		if err != nil {
			log.Printf("Error editing message: %v", err)
		}
	}
}
