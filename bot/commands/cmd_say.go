package commands

import (
	"gopkg.in/telegram-bot-api.v4"
)

func commandSay(bot *tgbotapi.BotAPI, args string, Context bool, update tgbotapi.Update) (err error) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, args)
	deleteMessageConfig := tgbotapi.DeleteMessageConfig{
		ChatID:    update.Message.Chat.ID,
		MessageID: update.Message.MessageID,
	}
	bot.DeleteMessage(deleteMessageConfig)
	bot.Send(msg)
	return nil
}
