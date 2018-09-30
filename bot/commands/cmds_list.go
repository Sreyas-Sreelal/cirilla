package commands

//Init loads all commands

func Init() map[string]Command {
	return map[string]Command{
		"/say": {
			Function:    commandSay,
			Description: "Say as Cirilla",
			PassString:  true,
			Admin:       true,
		},
	}
}
