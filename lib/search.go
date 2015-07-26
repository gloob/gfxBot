package gfxBot

import (
	"fmt"
	"strings"
	"github.com/tucnak/telebot"
)

var (
)

func SearchImage(bot *telebot.Bot, msg telebot.Message) (err error) {

	img, err := DuckSearch(strings.TrimPrefix(msg.Text, "/gfx"))
	if err != nil {
		fmt.Println(err)
		return err
	}

	filename := fmt.Sprint("assets/", msg.ID, img.Ext)

	err = img.Save(filename)
	if err != nil {
		fmt.Println(err)
		return err
	}

	i, err := telebot.NewFile(filename)
	if (err != nil) {
		fmt.Println(err)
		return err
	}

	photo := telebot.Photo{Thumbnail: telebot.Thumbnail{File: i, Width: img.Width, Height: img.Height}, Caption: img.Caption}

	err = bot.SendPhoto(msg.Chat, &photo, &telebot.SendOptions{ ReplyTo: msg })
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
