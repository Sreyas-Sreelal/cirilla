package commands

import (
	"github.com/Sreyas-Sreelal/cirilla/imdb"
	"github.com/Sreyas-Sreelal/cirilla/types"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"strings"
)

func commandMinfo(config *types.Config, bot *tgbotapi.BotAPI, args []string, Context bool, update tgbotapi.Update) (err error) {
	if len(strings.TrimSpace(args[0])) == 0 {
		NoArgsMessage := tgbotapi.NewMessage(update.Message.Chat.ID, `Invalid usage please provide arguement for this command.For example:\n`+config.CommandPrefix+"```minfo Now you see me```")
		NoArgsMessage.ReplyToMessageID = update.Message.MessageID
		bot.Send(NoArgsMessage)
		return
	}

	mc := imdb.GetNewClient()
	var botmsg tgbotapi.MessageConfig

	minfo, err := mc.GetMovieInfo(args[0])
	if err != nil {
		log.Print(err)
		botmsg = tgbotapi.NewMessage(update.Message.Chat.ID, "Failed to fetch information about that movie.Try again with precise name")
		bot.Send(botmsg)
		return nil
	}

	photomsg := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, nil)
	photomsg.FileID = minfo.PosterLink
	photomsg.UseExisting = true

	InfoMessage := "*Name* : " + minfo.Name + "\n```\n" + minfo.Description + "```\n*Rating* : " + minfo.Rating
	botmsg = tgbotapi.NewMessage(update.Message.Chat.ID, InfoMessage)
	botmsg.ParseMode = "markdown"
	bot.Send(botmsg)
	bot.Send(photomsg)

	return nil
}
