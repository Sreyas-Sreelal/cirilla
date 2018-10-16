package bot

import (
	"github.com/Sreyas-Sreelal/cirilla/bot/commands"
	"github.com/Sreyas-Sreelal/cirilla/types"
	"gopkg.in/telegram-bot-api.v4"
	"log"
)

//Start bot
func Start(config *types.Config) {

	bot, err := tgbotapi.NewBotAPI(config.TelegramToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = config.Debug

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(int(^uint(0) >> 1))
	u.Timeout = config.Timeout

	updates, err := bot.GetUpdatesChan(u)
	cmds := commands.Init()
	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if len(update.Message.Text) > 0 {

			if update.Message.Text[0] == config.CommandPrefix[0] {
				go commands.ExecuteCommand(update, cmds, bot)
			}
		}

	}
}
