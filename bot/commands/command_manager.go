package commands

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
)

//Command structure to represent commands
type Command struct {
	Function    func(bot *tgbotapi.BotAPI) (err error)
	Description string
}

//Init loads all commands
func Init() map[string]Command {
	return map[string]Command{
		"/say": {
			Function:    commandSay,
			Description: "Say as Cirilla",
		},
	}
}

//ExecuteCommand executes command
func ExecuteCommand(CommandName string, Commands map[string]Command, bot *tgbotapi.BotAPI) {
	if cmd, ok := Commands[CommandName]; ok {
		err := cmd.Function(bot)
		if err != nil {
			log.Println("Command : ", CommandName, " Failed to execute")
		}
	} else {
		log.Println("Unknown Command : ", CommandName, " Failed to execute")
	}

}
