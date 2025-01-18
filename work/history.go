package work

import (
	"bytes"
	"fmt"
	"os"

	"github.com/go-telegram/bot/models"
)

// TODO описание и камментарии
func (c *Configuration) HistoryPhoto(addressSlice []string) ([]models.InputMedia, error) {
	var photosMessageSlice []models.InputMedia

	for _, address := range addressSlice {

		if address[len(address)-4:] == ".jpg" {
			fileDataPhoto, err := os.ReadFile(address)
			if err != nil {
				fmt.Println("Ошибка при чтении файла - ", address, err)
				return photosMessageSlice, err
			}
			media := &models.InputMediaPhoto{
				Media:           "attach://" + address,
				MediaAttachment: bytes.NewReader(fileDataPhoto),
				ParseMode:       models.ParseModeMarkdownV1,
			}

			photosMessageSlice = append(photosMessageSlice, media)
		}

	}

	return photosMessageSlice, nil
}
