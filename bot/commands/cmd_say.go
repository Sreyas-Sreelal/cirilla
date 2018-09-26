package commands

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
)

func commandSay(bot *tgbotapi.BotAPI) (err error) {
	log.Println("Command executed")
	return nil
}
