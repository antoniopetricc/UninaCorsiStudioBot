package handlers

import (
	"fmt"
	"main/consts"
	"main/database"
	"main/services"

	"github.com/GoBotApiOfficial/gobotapi"
	"github.com/GoBotApiOfficial/gobotapi/methods"
	"github.com/GoBotApiOfficial/gobotapi/types"
)

func buildSearchCourseReplyMarkup() *types.InlineKeyboardMarkup {
	return &types.InlineKeyboardMarkup{
		InlineKeyboard: [][]types.InlineKeyboardButton{
			{
				{
					Text:         "🔙 Torna indietro",
					CallbackData: "start",
				},
			},
		},
	}
}

func SearchCourses(client *gobotapi.Client, update types.CallbackQuery) {

	services.UpdateStatus(update.From.ID, consts.SearchCoursesStatus)

	messageText := "🔎 <b>Cerca un corso</b>\n\n✏️ Digita il nome del corso che vuoi trovare nel campo qui sotto.\n\n📌 Puoi scrivere anche solo una parte del nome, il bot ti mostrerà tutti i risultati corrispondenti."

	client.Invoke(&methods.EditMessageText{
		ChatID:      update.Message.Chat.ID,
		MessageID:   update.Message.MessageID,
		Text:        messageText,
		ParseMode:   "HTML",
		ReplyMarkup: buildSearchCourseReplyMarkup(),
	})

}

func SearchCoursesByQuery(client *gobotapi.Client, update types.Message) {
	userStatus, err := services.GetUserStatus(update.From.ID)
	if err != nil {
		return
	}

	if userStatus == consts.SearchCoursesStatus {
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

		courses, total, err := database.SearchCourses(update.Text, 0)
		if err != nil {
			services.UpdateStatus(update.From.ID, consts.StartStatus)
			client.Invoke(&methods.SendMessage{
				ChatID: update.From.ID,
				Text:   "Si è verificato un errore durante la ricerca dei corsi.", ReplyMarkup: retryInlineKeyboard,
			})
			return
		}

		if total == 0 {
			services.UpdateStatus(update.From.ID, consts.SearchCoursesStatus)
			client.Invoke(&methods.SendMessage{
				ChatID: update.From.ID,
				Text:   "Nessun corso trovato con quel nome. Riprova con un altro termine di ricerca.", ReplyMarkup: retryInlineKeyboard,
			})
			return
		}

		services.UpdateStatus(update.From.ID, consts.StartStatus)

		var keyboard [][]types.InlineKeyboardButton

		coursesMessage := fmt.Sprintf("Ho trovato %d corsi:\n\n", len(courses))
		for _, course := range courses {
			keyboard = append(keyboard, []types.InlineKeyboardButton{
				{
					Text:         course.Nome + " (" + course.Cod + ")",
					CallbackData: fmt.Sprintf("course_%s", course.Cod),
				},
			})
		}

		coursesMessage += fmt.Sprintf("\nTotale corsi trovati: %d", len(courses))

		keyboard = append(keyboard, []types.InlineKeyboardButton{{
			Text:         "🔎 Cerca ancora",
			CallbackData: "search_courses",
		},
			{
				Text:         "🔙 Torna indietro",
				CallbackData: "start",
			},
		})

		client.Invoke(&methods.SendMessage{
			ChatID: update.From.ID,
			Text:   coursesMessage,
			ReplyMarkup: &types.InlineKeyboardMarkup{
				InlineKeyboard: keyboard,
			},
		})
	}
}
