package main

import (
	"fmt"
	"os/user"
	"strings"
	"time"

	"github.com/gloob/gfxBot/lib"
	"github.com/tucnak/telebot"
)

var (
	// Global configuration object.
	globalConfig gfxBot.GlobalConfig
)

func main() {
	// Load main configuration.
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	configFile := usr.HomeDir + "/.config/gfxBot/config.toml"
	err = gfxBot.LoadConfig(configFile, &globalConfig)
	if err != nil {
		panic(err)
	}

	bot, err := telebot.NewBot(globalConfig.Token)
	if err != nil {
		panic(err)
	}

	messages := make(chan telebot.Message)
	bot.Listen(messages, 1*time.Second)

	for message := range messages {
		fmt.Println(message)
		if message.Text == "/start" {
			bot.SendMessage(message.Chat, "This is a Telegram bot for searching images into different services.", nil)
		}
		if strings.HasPrefix(message.Text, "/gfx") {
			bot.SendChatAction(message.Chat, telebot.Typing)
			gfxBot.SearchImage(bot, message)
		}
		if message.Text == "/help" {
			bot.SendMessage(message.Chat, "This is a Telegram bot for searching images into different services.", nil)
		}
	}
}
