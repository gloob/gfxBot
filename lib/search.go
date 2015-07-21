package gfxBot

import (
	"fmt"
	"strings"
	"github.com/tucnak/telebot"
)

var (
)

func SearchImage(bot *telebot.Bot, msg telebot.Message) (err error) {

	filename, caption := DuckSearch(strings.TrimPrefix(msg.Text, "/gfx"))
	fmt.Println(filename)
	fmt.Println(caption)

	img, err := telebot.NewFile(filename)
	if (err != nil) {
		fmt.Println(err)
		return err
	}

	if caption == "" {
		bot.SendMessage(msg.Chat, "No results", nil)
		return nil
	}

	photo := telebot.Photo{Thumbnail: telebot.Thumbnail{File: img, Width: 32, Height: 32}, Caption: caption}

	err = bot.SendPhoto(msg.Chat, &photo, &telebot.SendOptions{ ReplyTo: msg })

	if err != nil {
		return err
	}

	return nil
}
