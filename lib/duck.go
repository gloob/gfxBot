package gfxBot

import (
	"errors"
	"github.com/ajanicij/goduckgo/goduckgo"
	"strings"
)

type Duck struct {
	Service string
}

func NewDuck() *Duck {
	return &Duck{
		Service: "DuckDuckGo",
	}
}

func (d *Duck) Search(q string) (*Image, error) {
	if q == "" {
		return nil, errors.New("DuckSearch: called with an empty query string.")
	}

	message, err := goduckgo.Query(q)
	if err != nil {
		return nil, err
	}

	var url, caption string
	// First we tried to obtain main Image from article.
	if message != nil && message.Image != "" {
		url = strings.TrimSpace(message.Image)
		caption = message.Heading
		// If there's no image we try to obtain an image or set of images from RelatedTopics.
	} else {
		for _, t := range message.RelatedTopics {
			if !t.Icon.IsEmpty() && t.Icon.URL != "" {
				url = strings.TrimSpace(t.Icon.URL)
				caption = t.Text
				break
			}
		}
	}

	if url != "" {
		img, _ := NewImage(url, caption)
		return img, nil
	} else {
		return nil, errors.New("DuckSearch error.")
	}
}
