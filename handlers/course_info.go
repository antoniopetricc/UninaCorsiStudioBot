package handlers

import (
	"fmt"
	"log"
	"main/database"
	"strings"

	"github.com/GoBotApiOfficial/gobotapi"
	"github.com/GoBotApiOfficial/gobotapi/methods"
	"github.com/GoBotApiOfficial/gobotapi/types"
)

func buildCourseInfoReplyMarkup() *types.InlineKeyboardMarkup {
	return &types.InlineKeyboardMarkup{
		InlineKeyboard: [][]types.InlineKeyboardButton{
			{
				{
					Text:         "🏫 Torna ai corsi",
					CallbackData: "explore_courses",
				},
			},
			{
				{
					Text:         "🔙 Torna al menu principale",
					CallbackData: "start",
				},
			},
		},
	}
}

func CourseInfo(client *gobotapi.Client, update types.CallbackQuery) {

	if strings.HasPrefix(update.Data, "course_") {

		course, err := database.GetCourse(strings.TrimPrefix(update.Data, "course_"))
		if err != nil {
			log.Printf("Error fetching course: %v", err)
			client.Invoke(&methods.EditMessageText{
				ChatID:    update.Message.Chat.ID,
				MessageID: update.Message.MessageID,
				Text:      "❌ Corso non trovato.",
			})
			return
		}

		messageText := fmt.Sprintf(
			`📚 <b>%s</b> (%s)

🏫 <b>Dipartimento:</b> %s
🎓 <b>Tipo di corso:</b> %s
⏳ <b>Durata:</b> %d anni
👨‍🏫 <b>Coordinatore:</b> %s %s
📧 <b>Email:</b> %s
🏛 <b>Sedi:</b> %s
🌐 <b>Lingua:</b> %s`,
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
			ReplyMarkup: buildCourseInfoReplyMarkup(),
		})

		if err != nil {
			log.Printf("Error editing message: %v", err)
		}
	}
}
