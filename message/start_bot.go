package message

import (
	"fmt"

	"TG_simple_bot/work"
)

// StartMessage - функция, которая возвращает в консоль сообщение про бот.
func StartMessage(configuration *work.Configuration) string {
	// текст приветственного сообщения состоящей из текста и данных о названии проекта из файла конфигурации
	greetings := fmt.Sprintf("Приветствую! Вы запускаете Телеграмм Бот Школьного проекта %s.", configuration.ProjectName)

	warning := "Если в процессе запуска бота что-то пойдёт не так, то тут появится соответствующее сообщение."

	// собираем результирующее сообщение
	resultMessage := fmt.Sprintf("%s\n%s", greetings, warning)

	return resultMessage
}

// OkMessage - функция, которая возвращает сообщение о том, что бот успешно запущен
func OkMessage() string {
	return "Бот успешно запущен!"
}

// FirstMessage - функция, которая возвращает сообщение перед началом диалога с ботом
func FirstMessage(configuration *work.Configuration) string {
	// текст приветственного сообщения состоящей из текста и данных о названии проекта из файла конфигурации
	greetings := fmt.Sprintf("Приветствую! Вы находетесь в Телеграмм Боте Школьного проекта %s.", configuration.ProjectName)

	// текст сообщения со ссылкой на Вики страницу проекта
	url := fmt.Sprintf("Ссылка на Wiki страницу проекта -  %s.", configuration.ProjectWikiURL)

	// собираем результирующее сообщение
	firstMessage := fmt.Sprintf("%s\n%s", greetings, url)

	return firstMessage
}
