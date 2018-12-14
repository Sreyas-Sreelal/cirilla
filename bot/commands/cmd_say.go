package commands

import (
	"github.com/Sreyas-Sreelal/cirilla/types"
	"gopkg.in/telegram-bot-api.v4"
)

func commandSay(config *types.Config, bot *tgbotapi.BotAPI, args []string, Context bool, update tgbotapi.Update) (err error) {
	if len(args) == 0 {
		return
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, args[0])
	deleteMessageConfig := tgbotapi.DeleteMessageConfig{
		ChatID:    update.Message.Chat.ID,
		MessageID: update.Message.MessageID,
	}
	bot.DeleteMessage(deleteMessageConfig)
	bot.Send(msg)
	return nil
}
