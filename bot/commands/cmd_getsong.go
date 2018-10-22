package commands

import (
	"github.com/Sreyas-Sreelal/cirilla/types"
	"github.com/Sreyas-Sreelal/cirilla/youtubedl"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"os"
)

func commandGetSong(config *types.Config, bot *tgbotapi.BotAPI, args []string, Context bool, update tgbotapi.Update) (err error) {
	if len(args) == 0 {
		NoArgsMessage := tgbotapi.NewMessage(update.Message.Chat.ID, "Invalid usage please provide arguement for this command.For example:\n**"+config.CommandPrefix+"getsong despacito**")
		NoArgsMessage.ReplyToMessageID = update.Message.MessageID
		NoArgsMessage.ParseMode = "markdown"
		bot.Send(NoArgsMessage)
		return
	}

	options := youtubedl.YtOptions{
		Path:        config.YotubedlPath,
		TrackName:   args[0],
		AudioFormat: "mp3",
	}
	FileName, err := youtubedl.YtSearchByName(options)
	FileName = FileName + "." + options.AudioFormat

	if err != nil {
		log.Print(err)
		botmsg := tgbotapi.NewMessage(update.Message.Chat.ID, "Failed to get that song :( Please try again with another name")
		botmsg.ReplyToMessageID = update.Message.MessageID
		bot.Send(botmsg)
		return err
	}
	log.Printf("Uploadingg %s", FileName)

	audiomsg := tgbotapi.NewAudioUpload(update.Message.Chat.ID, FileName)
	audiomsg.ReplyToMessageID = update.Message.MessageID
	audiomsg.FileID = FileName
	audiomsg.UseExisting = false

	_, err = bot.Send(audiomsg)
	if err != nil {
		log.Print(err)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Failed to upload if this problem persists report to Sreyas aka SyS")
		bot.Send(msg)
		os.Remove(FileName)
		return err
	}

	os.Remove(FileName)
	return nil
}
