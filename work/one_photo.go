package work

import (
	"bytes"
	"fmt"
	"os"

	"github.com/go-telegram/bot/models"
)

// TODO описание и камментарии
func (c *Configuration) OnePhoto() []models.InputMedia {

	fileDataPhoto, err := os.ReadFile("data/one_photo/фото_1шт.jpg")
	if err != nil {
		fmt.Println(err)
	}

	text := fmt.Sprint("текст для фото тут!")

	media1 := &models.InputMediaPhoto{
		Media:                 "attach://фото_1шт.jpg",
		Caption:               text,
		MediaAttachment:       bytes.NewReader(fileDataPhoto),
		ShowCaptionAboveMedia: true, // true - сверху фото, false - снизу фото
		ParseMode:             models.ParseModeMarkdownV1,
	}

	photo := []models.InputMedia{
		media1,
	}

	return photo
}
