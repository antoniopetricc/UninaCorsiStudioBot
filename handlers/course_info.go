package handlers

import (
	"fmt"
	"log"
	"main/database"
	"main/models"
	"strconv"
	"strings"

	"github.com/GoBotApiOfficial/gobotapi"
	"github.com/GoBotApiOfficial/gobotapi/methods"
	"github.com/GoBotApiOfficial/gobotapi/types"
)

func buildCourseInfoReplyMarkup(course models.Course, currentPage int) *types.InlineKeyboardMarkup {
	return &types.InlineKeyboardMarkup{
		InlineKeyboard: [][]types.InlineKeyboardButton{
			{
				{
					Text:         "📖 In breve",
					CallbackData: fmt.Sprintf("course_desc_%s_%d", course.Cod, currentPage),
				},
			},
			{
				{
					Text:         "🎒 Insegnamenti",
					CallbackData: fmt.Sprintf("course_teachings_%s_%d", course.Cod, currentPage),
				},
			},
			{
				{
					Text:         "🔙 Torna ai corsi",
					CallbackData: fmt.Sprintf("page_%d", currentPage),
				},
			},
		},
	}
}

func CourseInfo(client *gobotapi.Client, update types.CallbackQuery) {

	if strings.HasPrefix(update.Data, "course_info_") {

		courseCod := strings.SplitN(update.Data, "_", 4)[2]
		currentPage, err := strconv.Atoi(strings.SplitN(update.Data, "_", 4)[3])
		if err != nil {
			currentPage = 0
		}
		course, err := database.GetCourse(courseCod)

		if err != nil {
			log.Printf("Error fetching course: %v", err)
			client.Invoke(&methods.EditMessageText{
				ChatID:    update.Message.Chat.ID,
				MessageID: update.Message.MessageID,
				Text:      "❌ Corso non trovato.",
			})
			return
		}

		messageText := fmt.Sprintf("📚 <b>%s</b> (%s)\n\n🏫 <b>Dipartimento:</b> %s\n🎓 <b>Tipo di corso:</b> %s\n⏳ <b>Durata:</b> %d anni\n👨‍🏫 <b>Coordinatore:</b> %s %s\n📧 <b>Email:</b> %s\n🏛 <b>Sedi:</b> %s\n🌐 <b>Lingua:</b> %s",
			course.Nome,
			course.Cod,
			course.Dipartimento.DipDes,
			course.TipoCorso,
			course.Durata,

			course.CoordinatoreNome,
			course.CoordinatoreCognome,
			course.CoordinatoreEmail,
			course.Sedi,
			course.LinguaIta,
		)

		_, err = client.Invoke(&methods.EditMessageText{
			ChatID:      update.Message.Chat.ID,
			MessageID:   update.Message.MessageID,
			Text:        messageText,
			ParseMode:   "HTML",
			ReplyMarkup: buildCourseInfoReplyMarkup(course, currentPage),
		})

		if err != nil {
			log.Printf("Error editing message: %v", err)
		}
	}
}
