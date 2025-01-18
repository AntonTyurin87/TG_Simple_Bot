package work

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/go-telegram/bot/models"
)

// TODO описание и камментарии
func (c *Configuration) BusinessCard() []models.InputMedia {

	fileDataPhoto, err := os.ReadFile("data/business_card/team_photo.jpg")
	if err != nil {
		fmt.Println(err)
	}

	projectName := fmt.Sprintf("Команда проекта *\"%s\":*", c.ProjectName)
	projectTeam := "       " + strings.Replace(c.ProjectTeam, ", ", "\n       ", -1) + "\n"
	projectBriefDescription := fmt.Sprintf("*Краткое описание:* _%s_\n", c.BriefDescription)
	// projectWiki := fmt.Sprintf("[Wiki проекта](%s)\n", c.ProjectWikiURL)

	text := fmt.Sprintf("%s\n%s\n%s\n", projectName, projectTeam, projectBriefDescription)

	media1 := &models.InputMediaPhoto{
		Media:                 "attach://V1.jpg",
		Caption:               text,
		MediaAttachment:       bytes.NewReader(fileDataPhoto),
		ShowCaptionAboveMedia: true,
		ParseMode:             models.ParseModeMarkdownV1,
	}

	businessCard := []models.InputMedia{
		media1,
	}

	return businessCard
}
