package main

// библиотеки, необходимые для работы бота
import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"TG_simple_bot/menu"
	"TG_simple_bot/message"
	"TG_simple_bot/work"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// Переменная для хранения конфигурации
var Configuration *work.Configuration

// main - функция, в которой происходит инициализация бота и запуск его работы
func main() {
	// err - переменная для обработки ошибки
	var err error

	Configuration, err = work.SetNewConfig()
	if err != nil {
		fmt.Println("Ошибка конфигурации бота - ", err)
		panic(err)
	}

	// токен для авторизации бота в Телеграмм
	botToken := Configuration.Token

	// выводит сообщение о работе бота в консоль
	fmt.Println(message.StartMessage(Configuration))

	// TODO Сделать описание
	opts := []bot.Option{
		bot.WithDefaultHandler(defaultHandler),
		bot.WithCallbackQueryDataHandler("button", bot.MatchTypePrefix, callbackHandler),
	}

	// запуск и подключение бота по токену
	bot, err := bot.New(botToken, opts...)
	if err != nil {
		fmt.Println("Ошибка запуска бота - ", err)
		os.Exit(1)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	bot.Start(ctx)
}

func callbackHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       false,
	})

	// Подготовка данных ждя ответов:

	// descriptionText - переменная для хранения текста описания проекта
	descriptionText, err := work.ReadDescriptionFile(work.FindFiles("data/project_description"))
	if err != nil {
		fmt.Println("Ошибка обработки файла описания проекта - ", err)
		panic(err)
	}

	// isDisabledWikilink - переменная для отключения отображения содержания ссылки на Wiki
	isDisabledWikilink := true

	// photoAddress - переменная для хранения адресов фотографий
	photoAddress, err := work.FindFiles("data/history")
	if err != nil {
		fmt.Println("Ошибка сбора адресов фотографий для истории проекта - ", err)
		panic(err)
	}

	historyPhotos, err := Configuration.HistoryPhoto(photoAddress)
	if err != nil {
		fmt.Println("Ошибка обработки фотографий для истории проекта - ", err)
		panic(err)
	}

	// TODO описание
	documents, _ := work.FindFiles("data/project_materials")
	if len(documents) == 0 {
		err = fmt.Errorf("не обнаружены документы проекта")
		panic(err)
	}

	switch update.CallbackQuery.Data {
	case "button_1":
		// TODO Обработка ошибок сообщений!
		b.SendMediaGroup(ctx, &bot.SendMediaGroupParams{
			ChatID: update.CallbackQuery.Message.Message.Chat.ID,
			Media:  Configuration.BusinessCard(),
		})
		returnMenu(ctx, b, update)

	case "button_2":
		// TODO Обработка ошибок сообщений!
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    update.CallbackQuery.Message.Message.Chat.ID,
			Text:      descriptionText,
			ParseMode: models.ParseModeMarkdownV1,
		})
		returnMenu(ctx, b, update)

	case "button_3":
		// TODO Обработка ошибок сообщений!
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    update.CallbackQuery.Message.Message.Chat.ID,
			Text:      fmt.Sprintf("[Перейти на Wiki проекта \"%s\".](%s)", Configuration.ProjectName, Configuration.ProjectWikiURL),
			ParseMode: models.ParseModeMarkdownV1,
			LinkPreviewOptions: &models.LinkPreviewOptions{
				IsDisabled: &isDisabledWikilink,
			},
		})

	case "button_4":
		// TODO Обработка ошибок сообщений!
		_, err = b.SendMediaGroup(ctx, &bot.SendMediaGroupParams{
			ChatID: update.CallbackQuery.Message.Message.Chat.ID,
			Media:  historyPhotos,
		})
		if err != nil {
			fmt.Println(err)
		}
		returnMenu(ctx, b, update)

	case "button_5":
		for _, document := range documents {
			// TODO Обработка ошибок сообщений!
			// TODO может перенести отсюда
			projectDocuments, err := Configuration.Document(document)
			if err != nil {
				fmt.Println("шибка получения документов проекта - ", err)
				panic(err)
			}

			b.SendDocument(ctx, &bot.SendDocumentParams{
				ChatID:    update.CallbackQuery.Message.Message.Chat.ID,
				Document:  projectDocuments,
				ParseMode: models.ParseModeMarkdownV1,
			})
		}
		returnMenu(ctx, b, update)

	default:
		returnMenu(ctx, b, update)
	}
}

func defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	kb := menu.StartMenu()

	// TODO Обработка ошибок сообщений!
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Меню проекта *\"" + Configuration.ProjectName + "\"*",
		ReplyMarkup: kb,
		ParseMode:   models.ParseModeMarkdownV1,
	})
}

// TODO описание
func returnMenu(ctx context.Context, b *bot.Bot, update *models.Update) {
	kb := menu.StartMenu()

	// TODO Обработка ошибок сообщений!
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.CallbackQuery.Message.Message.Chat.ID,
		Text:        "Меню проекта *\"" + Configuration.ProjectName + "\"*",
		ReplyMarkup: kb,
		ParseMode:   models.ParseModeMarkdownV1,
	})
}
