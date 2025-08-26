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

func buildExploreCourseReplyMarkup(courses []models.Course, currentPage, totalPages int) *types.InlineKeyboardMarkup {
	var keyboard [][]types.InlineKeyboardButton

	for _, course := range courses {
		keyboard = append(keyboard, []types.InlineKeyboardButton{
			{
				Text:         course.Nome + " (" + course.Cod + ")",
				CallbackData: fmt.Sprintf("course_info_%s_%d", course.Cod, currentPage),
			},
		})
	}

	var navRow []types.InlineKeyboardButton
	if currentPage > 0 {
		navRow = append(navRow, types.InlineKeyboardButton{
			Text:         "⬅️ Pagina precedente",
			CallbackData: fmt.Sprintf("page_%d", currentPage-1),
		})
	}
	if currentPage < totalPages-1 {
		navRow = append(navRow, types.InlineKeyboardButton{
			Text:         "Pagina successiva ➡️",
			CallbackData: fmt.Sprintf("page_%d", currentPage+1),
		})
	}
	if len(navRow) > 0 {
		keyboard = append(keyboard, navRow)
	}

	keyboard = append(keyboard, []types.InlineKeyboardButton{
		{
			Text:         "🔙 Torna indietro",
			CallbackData: "start",
		},
	})

	return &types.InlineKeyboardMarkup{InlineKeyboard: keyboard}
}

func ExploreCourses(client *gobotapi.Client, update types.CallbackQuery) {

	if update.Data != "explore_courses" && !strings.HasPrefix(update.Data, "page_") {
		return
	}

	page := 0
	if update.Data == "explore_courses" {
		page = 0
	} else if strings.HasPrefix(update.Data, "page_") {
		if p, err := strconv.Atoi(strings.TrimPrefix(update.Data, "page_")); err == nil {
			page = p
		}
	}

	courses, total, err := database.GetCourses(page)
	if err != nil {
		log.Println("❌ Errore durante il recupero dei corsi:", err)
		client.Invoke(&methods.AnswerCallbackQuery{
			CallbackQueryID: update.ID,
			Text:            "❌ Errore durante il recupero dei corsi.",
			ShowAlert:       true,
		})
		return
	}

	exploreCoursesMessage := fmt.Sprintf("📚 Ecco i corsi <b>disponibili</b> (pagina %d/%d)\n\nℹ️ Clicca su un corso per avere maggiori <b>informazioni</b>.", page+1, total)

	client.Invoke(&methods.EditMessageText{
		ChatID:    update.Message.Chat.ID,
		MessageID: update.Message.MessageID,
		Text:      exploreCoursesMessage,
		ParseMode: "HTML",
		ReplyMarkup: buildExploreCourseReplyMarkup(
			courses,
			page,
			int(total),
		),
	})

}
