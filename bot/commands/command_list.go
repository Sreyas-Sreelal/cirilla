package commands

//Init loads all commands
func Init() map[string]Command {
	return map[string]Command{
		"say": {
			Function:    commandSay,
			Description: "Say as Cirilla",
			PassString:  true,
			Admin:       true,
		},
		"minfo": {
			Function:    commandMinfo,
			Description: "Get movie information",
			PassString:  true,
			Admin:       false,
		},
		"torrent": {
			Function:    commandTorrent,
			Description: "Get magnet Link of torrent",
			PassString:  true,
			Admin:       false,
		},
		"getsong": {
			Function:    commandGetSong,
			Description: "Get songs by name",
			PassString:  true,
			Admin:       false,
		},
		"yt2mp3": {
			Function:    commandYt2Mp3,
			Description: "Convert youtube videos to mp3",
			PassString:  true,
			Admin:       false,
		},
	}
}
