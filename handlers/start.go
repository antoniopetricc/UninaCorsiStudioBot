package handlers

import (
	"fmt"
	"html"
	"main/consts"
	"main/services"

	"github.com/GoBotApiOfficial/gobotapi"
	"github.com/GoBotApiOfficial/gobotapi/methods"
	"github.com/GoBotApiOfficial/gobotapi/types"
)

func buildStartKeyboard() *types.InlineKeyboardMarkup {
	return &types.InlineKeyboardMarkup{
		InlineKeyboard: [][]types.InlineKeyboardButton{
			{
				{Text: "🏫 Esplora Corsi", CallbackData: "explore_courses"},
			},
			{
				{Text: "🔎 Cerca per nome", CallbackData: "search_courses"},
			},
			{
				{Text: "🏛 Dipartimenti", CallbackData: "departments"},
			},
		},
	}
}

func buildWelcomeMessage(update types.Message) string {

	var fullName = html.EscapeString(update.From.FirstName + " " + update.From.LastName)
	mention := fmt.Sprintf(`<a href="tg://user?id=%d">%s</a>`, update.From.ID, fullName)

	return fmt.Sprintf(
		"👋 <b>Benvenuto</b> %s!\n\nBenvenuto nel bot per consultare i <b>Corsi di Studio</b> dell'<b>Università degli Studi di Napoli Federico II</b> 🎓\n\n📚 Usa i bottoni qui <b>sotto</b> per iniziare a esplorare i corsi disponibili.",
		mention,
	)
}

func Start(client *gobotapi.Client, update types.Message) {

	services.UpdateStatus(update.From.ID, consts.StartStatus)

	client.Invoke(&methods.SendMessage{
		ChatID:      update.Chat.ID,
		Text:        buildWelcomeMessage(update),
		ReplyMarkup: buildStartKeyboard(),
		ParseMode:   "HTML",
	})

}
