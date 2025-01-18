package work

import (
	"fmt"
	"os"

	"github.com/go-telegram/bot/models"
)

func (c *Configuration) Document(address string) (*models.InputFileUpload, error) {
	var fileData models.InputFileUpload

	file, err := os.Open(address)
	if err != nil {
		fmt.Println("Ошибка при открытии файла материалов проекта - ", address, err)
		return nil, err
	}

	fileData.Filename = address
	fileData.Data = file

	return &fileData, nil
}
