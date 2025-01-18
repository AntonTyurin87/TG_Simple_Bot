package menu

import (
	"github.com/go-telegram/bot/models"
)

// TODO описание и камментарии
func StartMenu() *models.InlineKeyboardMarkup {
	return &models.InlineKeyboardMarkup{InlineKeyboard: [][]models.InlineKeyboardButton{
		{
			{Text: "Визитка проекта", CallbackData: "button_1"},
		}, {
			{Text: "Описание проекта", CallbackData: "button_2"},
		}, {
			{Text: "Ссылка на Wiki проекта", CallbackData: "button_3"},
		}, {
			{Text: "История разработки проекта", CallbackData: "button_4"},
		}, {
			{Text: "Материалы проекта", CallbackData: "button_5"},
		},
	}}
}
