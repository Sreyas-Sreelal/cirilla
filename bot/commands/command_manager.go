package commands

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"strings"
)

//Command structure to represent commands
type Command struct {
	Function    func(bot *tgbotapi.BotAPI, args string, Context bool, update tgbotapi.Update) (err error)
	Description string
	Context     bool
}

//Init loads all commands
func Init() map[string]Command {
	return map[string]Command{
		"/say": {
			Function:    commandSay,
			Description: "Say as Cirilla",
			Context:     true,
		},
	}
}

//ExecuteCommand executes command
func ExecuteCommand(update tgbotapi.Update, Commands map[string]Command, bot *tgbotapi.BotAPI) {
	var CommandName string

	MessageSplitted := strings.SplitN(update.Message.Text, " ", 2)
	CommandName, args := MessageSplitted[0], MessageSplitted[1]
	if cmd, ok := Commands[CommandName]; ok {

		err := cmd.Function(bot, args, cmd.Context, update)
		if err != nil {
			log.Println("Command : ", CommandName, " Failed to execute")
		}
	} else {
		log.Println("Unknown Command : ", CommandName, " Failed to execute")
	}

}
