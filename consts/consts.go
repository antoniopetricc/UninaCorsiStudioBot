package consts

import (
	"reflect"
	"time"

	"github.com/GoBotApiOfficial/gobotapi/filters"
	"github.com/GoBotApiOfficial/gobotapi/types"
)

var (
	DefaultAntiFlood    = filters.AntiFlood(5, time.Second*5, time.Second*10)
	AliasList           = []string{"/", ";", ".", "!"}
	SearchCoursesStatus = "search_courses"
	StartStatus         = "start"
)

func Data(data string) filters.FilterOperand {
	return func(options *filters.DataFilter) bool {
		if reflect.TypeOf(options.RawUpdate).String() == "types.CallbackQuery" {
			val := options.RawUpdate.(types.CallbackQuery).Data
			return val == data
		}
		return false
	}
}
