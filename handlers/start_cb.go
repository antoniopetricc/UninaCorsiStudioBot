package handlers

import (
	"fmt"
	"main/consts"
	"main/misc"
	"main/services"

	"github.com/GoBotApiOfficial/gobotapi"
	"github.com/GoBotApiOfficial/gobotapi/methods"
	"github.com/GoBotApiOfficial/gobotapi/types"
)

func buildWelcomeCbMessage(update types.CallbackQuery) string {

	mention := misc.Mention(update.From.ID, fmt.Sprintf("%s %s", update.From.FirstName, update.From.LastName))

	return fmt.Sprintf(
		"👋 <b>Benvenuto</b> %s!\n\nBenvenuto nel bot per consultare i <b>Corsi di Studio</b> dell'<b>Università degli Studi di Napoli Federico II</b> 🎓\n\n📚 Usa i bottoni qui <b>sotto</b> per iniziare a esplorare i corsi disponibili.",
		mention,
	)
}

func StartCb(client *gobotapi.Client, update types.CallbackQuery) {

	services.UpdateStatus(update.From.ID, consts.StartStatus)

	client.Invoke(&methods.EditMessageText{
		ChatID:      update.Message.Chat.ID,
		MessageID:   update.Message.MessageID,
		Text:        buildWelcomeCbMessage(update),
		ReplyMarkup: buildStartKeyboard(),
		ParseMode:   "HTML",
	})

}
