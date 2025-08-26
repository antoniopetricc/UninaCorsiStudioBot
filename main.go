package main

import (
	"main/config"
	"main/consts"
	"main/handlers"
	"main/misc"

	"github.com/GoBotApiOfficial/gobotapi"
	"github.com/GoBotApiOfficial/gobotapi/filters"
	"github.com/GoBotApiOfficial/gobotapi/methods"
	"github.com/GoBotApiOfficial/gobotapi/types"
)

func main() {

	config.InitDatabase()
	config.LoadEnv()

	if config.GetEnv("LOAD_DEPARTMENTS") == "true" {
		misc.LoadDepartments()
	}
	if config.GetEnv("LOAD_COURSES") == "true" {
		misc.LoadCourses()
	}

	client := gobotapi.NewClient(config.GetEnv("BOT_API_TOKEN"))

	client.OnAnyMessageEvent(filters.Filter(handlers.Start, filters.And(filters.Command("start", consts.AliasList...), consts.DefaultAntiFlood)))

	client.OnCallbackQuery(func(client *gobotapi.Client, update types.CallbackQuery) {
		client.Invoke(&methods.AnswerCallbackQuery{
			CallbackQueryID: update.ID,
			Text:            "🔄 Caricamento...",
			ShowAlert:       false,
		})
	})

	client.OnCallbackQuery(filters.Filter(handlers.ExploreCourses, filters.And(consts.DefaultAntiFlood)))
	client.OnCallbackQuery(filters.Filter(handlers.SearchCourses, filters.And(consts.Data("search_courses"), consts.DefaultAntiFlood)))
	client.OnCallbackQuery(filters.Filter(handlers.Departments, filters.And(consts.Data("departments"), consts.DefaultAntiFlood)))
	client.OnCallbackQuery(filters.Filter(handlers.StartCb, filters.And(consts.Data("start"), consts.DefaultAntiFlood)))
	client.OnCallbackQuery(filters.Filter(handlers.CourseInfo, filters.And(consts.DefaultAntiFlood)))
	client.OnCallbackQuery(filters.Filter(handlers.CourseInfoDescription, filters.And(consts.DefaultAntiFlood)))
	client.OnAnyMessageEvent(filters.Filter(handlers.SearchCoursesByQuery, consts.DefaultAntiFlood))

	client.Run()

}
