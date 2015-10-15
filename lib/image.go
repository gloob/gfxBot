package gfxBot

import (
	"errors"
	"fmt"
	"github.com/tucnak/telebot"
	"math"
	"os"
)

const (
	MaxCaption = 255
)

type Image struct {
	Data     []byte
	Width    int
	Height   int
	Caption  string
	Ext      string
	Filename string
}

func NewImage(ext string, caption string) (*Image, error) {
	return &Image{Ext: ext, Caption: caption, Filename: ""}, nil
}

func (img *Image) Save(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(img.Data)
	if err != nil {
		return err
	}

	img.Filename = path

	return nil
}

func (img *Image) Write(p []byte) (n int, err error) {
	if img.Data == nil {
		img.Data = make([]byte, len(p))
		copy(img.Data, p)
	} else {
		img.Data = append(img.Data, p...)
	}
	return len(p), nil
}

func (img *Image) Close() (err error) {
	return
}

func (img *Image) Send(bot *telebot.Bot, msg telebot.Message) (err error) {
	if img == nil {
		warning := "You are trying to call a method of inexistent object :-)"
		bot.SendMessage(msg.Chat, warning, nil)
		return errors.New(warning)
	}
	img.Filename = fmt.Sprint("assets/", msg.ID, img.Ext)
	if img.Filename == "" {
		bot.SendMessage(msg.Chat, "There's any filename associated to this query.", nil)
		return errors.New("There's any filename associated to this query.")
	}

	i, err := telebot.NewFile(img.Filename)
	if err != nil {
		return err
	}

	caption := img.Caption[:int(math.Min(float64(len(img.Caption)), MaxCaption))]
	photo := telebot.Photo{Thumbnail: telebot.Thumbnail{File: i, Width: img.Width, Height: img.Height}, Caption: caption}

	err = bot.SendPhoto(msg.Chat, &photo, &telebot.SendOptions{ReplyTo: msg})
	if err != nil {
		return err
	}

	return nil
}
