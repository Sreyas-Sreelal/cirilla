package commands

import (
	"github.com/Sreyas-Sreelal/cirilla/torrent"
	"github.com/Sreyas-Sreelal/cirilla/types"
	"gopkg.in/telegram-bot-api.v4"
	"log"
)

func commandTorrent(config *types.Config, bot *tgbotapi.BotAPI, args []string, Context bool, update tgbotapi.Update) (err error) {
	pb := torrent.GetNewClient()
	var botmsg tgbotapi.MessageConfig

	torrentInfo, err := pb.GetTorrentInfo(args[0])
	if err != nil {
		log.Print(err)
		botmsg = tgbotapi.NewMessage(update.Message.Chat.ID, "Failed to fetch information about that torrent.Try again with precise name")
		bot.Send(botmsg)
		return nil
	}

	InfoMessage := "Name : " + torrentInfo.Name + "\n" + torrentInfo.Description + "\nLink : " + torrentInfo.URL
	MagnetMessage := torrentInfo.MagnetURL

	botmsg = tgbotapi.NewMessage(update.Message.Chat.ID, InfoMessage)
	bot.Send(botmsg)

	botmsg = tgbotapi.NewMessage(update.Message.Chat.ID, MagnetMessage)
	bot.Send(botmsg)

	return nil
}
