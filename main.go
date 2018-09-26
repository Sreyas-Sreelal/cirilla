package main

import (
	"github.com/Sreyas-Sreelal/cirilla/bot"
	"github.com/Sreyas-Sreelal/cirilla/types"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	config := &types.Config{}
	err := envconfig.Process("CIRILLA", config)
	if err != nil {
		panic(err)
	}
	bot.Start(config)

}
