package work

import (
	"encoding/json"
	"fmt"
	"os"
)

// Configuration - структура для работы с конфигурациями
type Configuration struct {
	Token            string `json:"token"`
	ProjectName      string `json:"project_name"`
	ProjectWikiURL   string `json:"project_wiki_url"`
	BriefDescription string `json:"brief_description"`
	ProjectTeam      string `json:"project_team"`
}

// SetNewConfig - функция для чтения файла конфигурации, возвращает структуру с конфигурацией
func SetNewConfig() (*Configuration, error) {
	// Читаем файл с конфигурацией, если файл не читается, то сообщаем об ошибке
	configFile, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Println("Ошибка открытия конфиг файла - ", err)
		return nil, err
	}

	// botConfiguration - переменная, в которую мы положим информацию о конфигурации бота
	var botConfiguration Configuration

	// распаковка из данных файла в данные структуры конфигурации
	err = json.Unmarshal(configFile, &botConfiguration)
	if err != nil {
		fmt.Println("Ошибка распаковки json файла - ", err)
		return nil, err
	}

	return &botConfiguration, nil // TODO Обработка ошибки
}
