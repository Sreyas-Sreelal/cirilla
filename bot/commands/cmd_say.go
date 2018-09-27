package commands

import (
	"gopkg.in/telegram-bot-api.v4"
)

func commandSay(bot *tgbotapi.BotAPI, args string, Context bool, update tgbotapi.Update) (err error) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, args)
	bot.Send(msg)
	return nil
}
