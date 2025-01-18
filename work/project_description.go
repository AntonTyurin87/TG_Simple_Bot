package work

import (
	"fmt"
	"os"
)

// TODO сделать описание и комментарии
func FindFiles(folder string) ([]string, error) {
	fileList := []string{}

	// Открываем директорию
	dir, err := os.Open(folder)
	if err != nil {
		fmt.Println("Ошибка при открытии директории с описанием проекта - ", err)
		return fileList, err
	}
	defer dir.Close()

	// Получаем список файлов и папок
	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("Ошибка при получении списка файлов для описания проекта - ", err)
		return fileList, err
	}

	for _, file := range files {
		fileList = append(fileList, folder+"/"+file.Name())
	}

	return fileList, nil
}

// TODO сделать описание и комментарии
func ReadDescriptionFile(fileList []string, err error) (string, error) {
	// TODO Вот тут сообщить об ошибке поиска имени файла и выше переделать
	if len(fileList) != 1 {
		err := fmt.Errorf("количество файлов для описания проекта отличается от одного")
		return "", err
	}

	if fileList[0][len(fileList[0])-4:] != ".txt" {
		err := fmt.Errorf("формат файла описания проекта отличается от \".txt\"")
		return "", err
	}

	file, err := os.ReadFile(fileList[0])
	if err != nil {
		fmt.Println("Не удалось прочитать файл описания проекта - ", err)
		return "", err
	}

	projectDescription := string(file)

	return projectDescription, nil
}
