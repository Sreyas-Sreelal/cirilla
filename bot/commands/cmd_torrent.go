package commands

import (
	"github.com/Sreyas-Sreelal/cirilla/torrent"
	"github.com/Sreyas-Sreelal/cirilla/types"
	"gopkg.in/telegram-bot-api.v4"
	"log"
)

func commandTorrent(config *types.Config, bot *tgbotapi.BotAPI, args []string, Context bool, update tgbotapi.Update) (err error) {
	if len(args) == 0 {
		NoArgsMessage := tgbotapi.NewMessage(update.Message.Chat.ID, "Invalid usage please provide arguement for this command.For example:\n**"+config.CommandPrefix+"torrent avengers infinity war**")
		NoArgsMessage.ReplyToMessageID = update.Message.MessageID
		NoArgsMessage.ParseMode = "markdown"
		bot.Send(NoArgsMessage)
		return
	}

	pb := torrent.GetNewClient()
	var botmsg tgbotapi.MessageConfig

	botmsg.ReplyToMessageID = update.Message.MessageID
	torrentInfo, err := pb.GetTorrentInfo(args[0])
	if err != nil {
		log.Print(err)
		botmsg = tgbotapi.NewMessage(update.Message.Chat.ID, "Failed to fetch information about that torrent.Try again with precise name")
		bot.Send(botmsg)
		return nil
	}

	InfoMessage := "*Name* : " + torrentInfo.Name + "\n```\n" + torrentInfo.Description + "```\n*Link* : " + torrentInfo.URL
	MagnetMessage := "```\n" + torrentInfo.MagnetURL + "```"

	botmsg = tgbotapi.NewMessage(update.Message.Chat.ID, InfoMessage)
	botmsg.ParseMode = "markdown"
	bot.Send(botmsg)

	botmsg = tgbotapi.NewMessage(update.Message.Chat.ID, MagnetMessage)
	botmsg.ParseMode = "markdown"
	bot.Send(botmsg)

	return nil
}
