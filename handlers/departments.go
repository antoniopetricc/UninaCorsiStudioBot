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

func buildDepartmentsReplyMarkup() *types.InlineKeyboardMarkup {
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

func Departments(client *gobotapi.Client, update types.CallbackQuery) {

	departments, err := database.GetDepartments()
	if err != nil {
		client.Invoke(&methods.AnswerCallbackQuery{
			CallbackQueryID: update.ID,
			Text:            "❌ Errore durante il recupero dei dipartimenti	.",
			ShowAlert:       true,
		})
		return
	}

	messageText := "<b>🏛 Dipartimenti disponibili</b>\n\n"
	for i, dept := range departments {
		if dept.Cod == "" {
			continue
		}
		messageText += fmt.Sprintf("• %s\n", dept.DipDes)
		if i == len(departments)-1 {
			messageText += "\n"
		}
	}

	services.UpdateStatus(update.From.ID, consts.StartStatus)

	client.Invoke(&methods.EditMessageText{
		ChatID:      update.Message.Chat.ID,
		MessageID:   update.Message.MessageID,
		Text:        messageText,
		ParseMode:   "HTML",
		ReplyMarkup: buildDepartmentsReplyMarkup(),
	})

}
