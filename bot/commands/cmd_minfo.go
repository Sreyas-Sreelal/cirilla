package commands

import (
	"github.com/Sreyas-Sreelal/cirilla/imdb"
	"github.com/Sreyas-Sreelal/cirilla/types"
	"gopkg.in/telegram-bot-api.v4"
	"log"
)

func commandMinfo(config *types.Config, bot *tgbotapi.BotAPI, args []string, Context bool, update tgbotapi.Update) (err error) {
	if len(args) == 0 {
		NoArgsMessage := tgbotapi.NewMessage(update.Message.Chat.ID, "Invalid usage please provide arguement for this command.For example:\n**"+config.CommandPrefix+"minfo Now you see me**")
		NoArgsMessage.ReplyToMessageID = update.Message.MessageID
		NoArgsMessage.ParseMode = "markdown"
		bot.Send(NoArgsMessage)
		return
	}

	mc := imdb.GetNewClient()

	minfo, err := mc.GetMovieInfo(args[0])
	if err != nil {
		log.Print(err)
		botmsg := tgbotapi.NewMessage(update.Message.Chat.ID, "Failed to fetch information about that movie.Try again with precise name")
		botmsg.ReplyToMessageID = update.Message.MessageID
		bot.Send(botmsg)
		return nil
	}

	InfoMessage := "Name : " + minfo.Name + "\n" + minfo.Description + "\nRating : " + minfo.Rating
	photomsg := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, nil)
	photomsg.FileID = minfo.PosterLink
	photomsg.UseExisting = true
	photomsg.Caption = InfoMessage
	photomsg.ReplyToMessageID = update.Message.MessageID
	bot.Send(photomsg)
	return nil
}
