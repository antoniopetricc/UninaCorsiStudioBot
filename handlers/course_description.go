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

func buildCourseInfoDescriptionReplyMarkup(course models.Course, currentPage int) *types.InlineKeyboardMarkup {
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

func CourseInfoDescription(client *gobotapi.Client, update types.CallbackQuery) {

	if strings.HasPrefix(update.Data, "course_desc_") {

		// course_desc_<cod>_<page>

		courseCod := strings.SplitN(update.Data, "_", 4)[2]
		currentPage, err := strconv.Atoi(strings.SplitN(update.Data, "_", 4)[3])
		if err != nil {
			currentPage = 0
		}

		course, err := database.GetCourse(courseCod)

		if course.InBreve == "" {
			courseDescriptionResponse, err := services.GetCourseDescription(course.Cod)
			if err != nil {
				log.Printf("Error fetching course description: %v", err)
				course.InBreve = "Descrizione non disponibile."
			} else {
				course.InBreve = courseDescriptionResponse.Content
				err = database.UpdateCourseDescription(course.Cod, course.InBreve)
				if err != nil {
					log.Printf("Error updating course description: %v", err)
				}
			}
		}

		if err != nil {
			retryInlineKeyboard := &types.InlineKeyboardMarkup{
				InlineKeyboard: [][]types.InlineKeyboardButton{
					{
						{
							Text:         "🔎 Cerca ancora",
							CallbackData: "search_courses",
						},
					},
					{
						{
							Text:         "🔙 Torna indietro",
							CallbackData: "start",
						},
					},
				},
			}
			log.Printf("Error fetching course: %v", err)
			client.Invoke(&methods.EditMessageText{
				ChatID:      update.Message.Chat.ID,
				MessageID:   update.Message.MessageID,
				Text:        "❌ Corso non trovato.",
				ReplyMarkup: retryInlineKeyboard,
			})
			return
		}

		messageText := fmt.Sprintf("ℹ️ <b>In breve su %s (%s)</b>\n\n%s",
			course.Nome,
			course.Cod,
			course.InBreve,
		)

		_, err = client.Invoke(&methods.EditMessageText{
			ChatID:      update.Message.Chat.ID,
			MessageID:   update.Message.MessageID,
			Text:        messageText,
			ParseMode:   "HTML",
			ReplyMarkup: buildCourseInfoDescriptionReplyMarkup(course, currentPage),
		})

		if err != nil {
			log.Printf("Error editing message: %v", err)
		}
	}
}
