package main

import (
	"fmt"
	"log"
	"os/user"
	"strings"
	"time"

	"github.com/gloob/gfxBot/lib"
	"github.com/stathat/go"
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
		log.Fatal(err)
	}
	fmt.Println(usr)
	//configFile := usr.HomeDir + "/.config/gfxbot/config.toml"
	configFile := "config.toml"
	err = gfxBot.LoadConfig(configFile, &globalConfig)
	if err != nil {
		log.Fatal(err)
	}

	bot, err := telebot.NewBot(globalConfig.Token)
	if err != nil {
		log.Fatal(err)
	}

	messages := make(chan telebot.Message)
	bot.Listen(messages, 1*time.Second)

	for message := range messages {
		fmt.Println(message)
		stathat.PostEZCount("gfxBot message processed", "gloob@litio.org", 1)
		if message.Text == "/start" {
			bot.SendMessage(message.Chat, "This is a Telegram bot for searching images into different services.", nil)
		}
		if strings.HasPrefix(message.Text, "/duck") {
			bot.SendChatAction(message.Chat, telebot.Typing)
			duck := gfxBot.NewDuck()
			img, err := duck.Search(strings.TrimPrefix(strings.TrimPrefix(message.Text, "/duck"), "@gfxBot"))
			if err != nil {
				fmt.Println("Error in duck.Search()")
				fmt.Println(err)
				bot.SendMessage(message.Chat, fmt.Sprintf("%s", err), nil)
				continue
			}
			filename := fmt.Sprint("assets/", message.ID, img.Ext)
			err = img.Save(filename)
			err = img.Send(bot, message)
			if err != nil {
				bot.SendMessage(message.Chat, fmt.Sprintf("%s", err), nil)
			}
		}
		if strings.HasPrefix(message.Text, "/flickr") {
			bot.SendChatAction(message.Chat, telebot.Typing)
			flickr := gfxBot.NewFlickr(globalConfig.FlickrAPIKey, globalConfig.FlickrAPISecret)
			img, err := flickr.Search(strings.TrimPrefix(strings.TrimPrefix(message.Text, "/flickr"), "@gfxBot"))
			if err != nil {
				fmt.Println("Error in flickr.Search()")
				fmt.Println(err)
				bot.SendMessage(message.Chat, fmt.Sprintf("%s", err), nil)
				continue
			}
			filename := fmt.Sprint("assets/", message.ID, img.Ext)
			err = img.Save(filename)
			err = img.Send(bot, message)
			if err != nil {
				bot.SendMessage(message.Chat, fmt.Sprintf("%s", err), nil)
			}
		}
		if message.Text == "/help" {
			bot.SendMessage(message.Chat, "This is a Telegram bot for searching images into different services.", nil)
		}
	}
}
