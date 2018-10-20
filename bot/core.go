package bot

import (
	"github.com/Sreyas-Sreelal/cirilla/bot/commands"
	"github.com/Sreyas-Sreelal/cirilla/types"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"time"
)

//Start bot
func Start(config *types.Config) {

	StartedTimeStamp := time.Now()
	bot, err := tgbotapi.NewBotAPI(config.TelegramToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = config.Debug

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = config.Timeout

	updates, err := bot.GetUpdatesChan(u)
	cmds := commands.Init()
	for update := range updates {

		if update.Message == nil || update.Message.Time().Before(StartedTimeStamp) {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if len(update.Message.Text) > 0 {

			if update.Message.Text[0] == config.CommandPrefix[0] {
				go commands.ExecuteCommand(config, update, cmds, bot)
			}
		}

	}
}
