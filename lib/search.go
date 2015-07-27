package gfxBot

import (
	"fmt"
	"strings"
	"github.com/tucnak/telebot"
)

type Searcher interface {
	Search(q string) (*Image, error)
}

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

	err = img.Send(bot, msg)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
